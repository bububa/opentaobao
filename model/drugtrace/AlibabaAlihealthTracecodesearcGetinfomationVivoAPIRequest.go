package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest 获取vivo banner API请求
// alibaba.alihealth.tracecodesearc.getinfomation.vivo
//
// 获取vivo banner  url
type AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest struct {
	model.Params
	// 渠道
	_channel string
}

// NewAlibabaalihealthtracecodesearcgetinfomationvivoRequest 初始化AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest对象
func NewAlibabaalihealthtracecodesearcgetinfomationvivoRequest() *AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest {
	return &AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodesearc.getinfomation.vivo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannel is Channel Setter
// 渠道
func (r *AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest) GetChannel() string {
	return r._channel
}
