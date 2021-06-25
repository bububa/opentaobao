package core

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/bububa/opentaobao/core/internal/debug"
	"github.com/bububa/opentaobao/model"
)

type SDKClient struct {
	appKey  string
	secret  string
	debug   bool
	sandbox bool
}

func NewSDKClient(appKey string, secret string) *SDKClient {
	return &SDKClient{
		appKey: appKey,
		secret: secret,
	}
}

func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

func (c *SDKClient) Post(req model.IRequest, resp model.IResponse, session string) error {
	commonReq := model.NewCommonRequest(req.GetApiMethodName(), c.appKey)
	commonReq.SetSession(session)
	params := c.sign(commonReq, req)

	var err error
	if req.NeedMultipart() {
		err = c.postMultipart(params, req.GetRawParams(), resp)
	} else {
		err = c.post(params, resp)
	}
	return err
}

func (c *SDKClient) post(req url.Values, resp model.IResponse) error {
	reqUrl := PRODUCT_GATEWAY
	if c.sandbox {
		reqUrl = SANDBOX_GATEWAY
	}
	debug.PrintPostJSONRequest(reqUrl, req.Encode(), c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, strings.NewReader(req.Encode()))
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

func (c *SDKClient) postMultipart(req url.Values, params model.Params, resp model.IResponse) error {
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	for k, _ := range req {
		var (
			fw  io.Writer
			r   io.Reader
			err error
		)

		if v, found := params[k]; found && v.IsFile() {
			if fw, err = mw.CreateFormFile(k, v.String()); err != nil {
				return err
			}
			r = v.File()
		} else {
			if fw, err = mw.CreateFormField(k); err != nil {
				return err
			}
			r = strings.NewReader(req.Get(k))
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	reqUrl := PRODUCT_GATEWAY
	if c.sandbox {
		reqUrl = SANDBOX_GATEWAY
	}
	debug.PrintPostJSONRequest(reqUrl, req.Encode(), c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, &buf)
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

func (c *SDKClient) get(req url.Values, resp model.IResponse) error {
	reqUrl := PRODUCT_GATEWAY
	if c.sandbox {
		reqUrl = SANDBOX_GATEWAY
	}
	reqUrl = fmt.Sprintf("%s?%s", reqUrl, req.Encode())
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

func (c *SDKClient) sign(commonReq *model.CommonRequest, req model.IRequest) url.Values {
	var ret url.Values
	commonParams := commonReq.GetParams()
	reqParams := req.GetApiParams()
	keys := make([]string, 0, len(commonParams)+len(reqParams))
	for k, v := range commonParams {
		keys = append(keys, k)
		ret.Set(k, v)
	}
	for k := range reqParams {
		keys = append(keys, k)
		ret.Set(k, reqParams.Get(k))
	}
	sort.Strings(keys)
	query := bytes.NewBufferString(c.secret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(ret.Get(k))
	}
	query.WriteString(c.secret)
	h := md5.New()
	io.Copy(h, query)
	sign := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	ret.Set("sign", sign)
	return ret
}
