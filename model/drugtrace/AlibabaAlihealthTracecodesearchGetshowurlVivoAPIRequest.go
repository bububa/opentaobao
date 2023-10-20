package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest 获取药品扫码落地页vivo API请求
// alibaba.alihealth.tracecodesearch.getshowurl.vivo
//
// 获取药品扫码落地页vivo
type AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 来源
	_channel string
}

// NewAlibabaalihealthtracecodesearchgetshowurlvivoRequest 初始化AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest对象
func NewAlibabaalihealthtracecodesearchgetshowurlvivoRequest() *AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest {
	return &AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodesearch.getshowurl.vivo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest) GetCode() string {
	return r._code
}

// SetChannel is Channel Setter
// 来源
func (r *AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabaalihealthtracecodesearchgetshowurlvivoAPIRequest) GetChannel() string {
	return r._channel
}
