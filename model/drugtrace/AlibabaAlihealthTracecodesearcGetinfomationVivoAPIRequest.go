package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest 获取vivo banner API请求
// alibaba.alihealth.tracecodesearc.getinfomation.vivo
//
// 获取vivo banner  url
type AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest struct {
	model.Params
	// 渠道
	_channel string
}

// NewAlibabaAlihealthTracecodesearcGetinfomationVivoRequest 初始化AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest对象
func NewAlibabaAlihealthTracecodesearcGetinfomationVivoRequest() *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest {
	return &AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodesearc.getinfomation.vivo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetChannel is Channel Setter
// 渠道
func (r *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) GetChannel() string {
	return r._channel
}
