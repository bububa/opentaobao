package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest 获取药品扫码落地页vivo API请求
// alibaba.alihealth.tracecodesearch.getshowurl.vivo
//
// 获取药品扫码落地页vivo
type AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 来源
	_channel string
}

// NewAlibabaAlihealthTracecodesearchGetshowurlVivoRequest 初始化AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest对象
func NewAlibabaAlihealthTracecodesearchGetshowurlVivoRequest() *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest {
	return &AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) Reset() {
	r._code = ""
	r._channel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodesearch.getshowurl.vivo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) GetCode() string {
	return r._code
}

// SetChannel is Channel Setter
// 来源
func (r *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) GetChannel() string {
	return r._channel
}

var poolAlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesearchGetshowurlVivoRequest()
	},
}

// GetAlibabaAlihealthTracecodesearchGetshowurlVivoRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest
func GetAlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest() *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest {
	return poolAlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest.Get().(*AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest 将 AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest(v *AlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesearchGetshowurlVivoAPIRequest.Put(v)
}
