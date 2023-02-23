package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscChudaTemplateSendAPIRequest 本地生活触达模板消息发送接口 API请求
// alibaba.alsc.chuda.template.send
//
// 允许三方小程序通过该api发送本地生活触达消息
type AlibabaAlscChudaTemplateSendAPIRequest struct {
	model.Params
	// 请求参数
	_notifyRequest *TemplateNotifyRequest
}

// NewAlibabaAlscChudaTemplateSendRequest 初始化AlibabaAlscChudaTemplateSendAPIRequest对象
func NewAlibabaAlscChudaTemplateSendRequest() *AlibabaAlscChudaTemplateSendAPIRequest {
	return &AlibabaAlscChudaTemplateSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscChudaTemplateSendAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.chuda.template.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscChudaTemplateSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscChudaTemplateSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNotifyRequest is NotifyRequest Setter
// 请求参数
func (r *AlibabaAlscChudaTemplateSendAPIRequest) SetNotifyRequest(_notifyRequest *TemplateNotifyRequest) error {
	r._notifyRequest = _notifyRequest
	r.Set("notify_request", _notifyRequest)
	return nil
}

// GetNotifyRequest NotifyRequest Getter
func (r AlibabaAlscChudaTemplateSendAPIRequest) GetNotifyRequest() *TemplateNotifyRequest {
	return r._notifyRequest
}
