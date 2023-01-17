package core

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/bububa/opentaobao/core/internal/debug"
	"github.com/bububa/opentaobao/metadata/util"
	"github.com/bububa/opentaobao/model"
)

// SDKClient 结构体
type SDKClient struct {
	appKey     string           // APP KEY
	secret     string           // APP SECRET
	apiFormat  model.APIFormat  // API 响应格式
	signMethod model.SignMethod // API 签名方法
	debug      bool             // debug
	sandbox    bool             // 是否沙箱环境
}

// NewSDKClient 新建SDKClient
func NewSDKClient(appKey string, secret string) *SDKClient {
	return &SDKClient{
		appKey:     appKey,
		secret:     secret,
		apiFormat:  model.DEFAULT_API_FORMAT,
		signMethod: model.DEFAULT_SIGN_METHOD,
	}
}

// SetDebug 设置debug
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// IsDebug 判断是否debug
func (c SDKClient) IsDebug() bool {
	return c.debug
}

// UseSandbox 使用沙箱环境
func (c *SDKClient) UseSandbox() {
	c.sandbox = true
}

// DisableSandbox 禁用沙箱环境
func (c *SDKClient) DisableSandbox() {
	c.sandbox = false
}

// IsSandbox 判断是否沙箱环境
func (c SDKClient) IsSandbox() bool {
	return c.sandbox
}

// SetAPIFormat 设置API响应格式
func (c *SDKClient) SetAPIFormat(format model.APIFormat) {
	c.apiFormat = format
}

// SetSignMethod 设置API签名方法
func (c *SDKClient) SetSignMethod(method model.SignMethod) {
	c.signMethod = method
}

// Post 发起请求
func (c *SDKClient) Post(req model.IRequest, resp model.IResponse, session string) error {
	// 新建API请求通用参数
	commonReq := model.NewCommonRequest(req.GetApiMethodName(), c.appKey)
	commonReq.SetSession(session)
	commonReq.SetAPIFormat(c.apiFormat)
	commonReq.SetSignMethod(c.signMethod)
	params := util.GetUrlValues()
	c.sign(params, commonReq, req)

	var err error
	if req.NeedMultipart() {
		err = c.postMultipart(params, req.GetRawParams(), resp)
	} else {
		err = c.post(params, resp)
	}
	util.PutUrlValues(params)
	return err
}

// post application/xml-www-form-urlencode
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
	switch c.apiFormat {
	case model.JSON:
		err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	case model.XML:
		err = debug.DecodeXMLHttpResponse(httpResp.Body, resp, c.debug)
	default:
		panic("unknown api format")
	}
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return resp.B043C16EB094F65A787F22E6AE0A10BCB7ABDE6D()
}

// postMultipart post multipart/form
func (c *SDKClient) postMultipart(req url.Values, params model.Params, resp model.IResponse) error {
	buf := util.GetBufferPool()
	defer util.PutBufferPool(buf)
	mw := multipart.NewWriter(buf)
	for k := range req {
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
	httpReq, err := http.NewRequest("POST", reqUrl, buf)
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	if err != nil {
		return err
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	switch c.apiFormat {
	case model.JSON:
		err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	case model.XML:
		err = debug.DecodeXMLHttpResponse(httpResp.Body, resp, c.debug)
	default:
		panic("unknown api format")
	}
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return resp.B043C16EB094F65A787F22E6AE0A10BCB7ABDE6D()
}

// get (not used)
func (c *SDKClient) get(req url.Values, resp model.IResponse) error {
	reqUrl := PRODUCT_GATEWAY
	if c.sandbox {
		reqUrl = SANDBOX_GATEWAY
	}
	reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
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
	switch c.apiFormat {
	case model.JSON:
		err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	case model.XML:
		err = debug.DecodeXMLHttpResponse(httpResp.Body, resp, c.debug)
	default:
		panic("unknown api format")
	}
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return resp.B043C16EB094F65A787F22E6AE0A10BCB7ABDE6D()
}

// sign 生成签名
// 支持md5, hmac
func (c *SDKClient) sign(ret url.Values, commonReq *model.CommonRequest, req model.IRequest) {
	commonReq.GetParams(ret)
	req.GetApiParams(ret)
	keys := make([]string, 0, len(ret))
	for k := range ret {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	query := util.GetBufferPool()
	if commonReq.SignMethod == model.MD5 {
		query.WriteString(c.secret)
	}
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(ret.Get(k))
	}
	if commonReq.SignMethod == model.MD5 {
		query.WriteString(c.secret)
	}
	var h hash.Hash
	switch commonReq.SignMethod {
	case model.MD5:
		h = md5.New()
	case model.HMAC:
		h = hmac.New(md5.New, []byte(c.secret))
	default:
		panic("missing sign_method")
	}
	io.Copy(h, query)
	util.PutBufferPool(query)
	sign := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	ret.Set("sign", sign)
}
