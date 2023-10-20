package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) Reset() {
	r._channel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodesearc.getinfomation.vivo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesearcGetinfomationVivoRequest()
	},
}

// GetAlibabaAlihealthTracecodesearcGetinfomationVivoRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest
func GetAlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest() *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest {
	return poolAlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest.Get().(*AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest 将 AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest(v *AlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesearcGetinfomationVivoAPIRequest.Put(v)
}
