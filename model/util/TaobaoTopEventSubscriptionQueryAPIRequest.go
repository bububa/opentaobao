package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopEventSubscriptionQueryAPIRequest 三方事件订阅查询 API请求
// taobao.top.event.subscription.query
//
// 三方事件订阅查询
type TaobaoTopEventSubscriptionQueryAPIRequest struct {
	model.Params
	// 连接平台触发事件TriggerCode
	_triggerCode string
}

// NewTaobaoTopEventSubscriptionQueryRequest 初始化TaobaoTopEventSubscriptionQueryAPIRequest对象
func NewTaobaoTopEventSubscriptionQueryRequest() *TaobaoTopEventSubscriptionQueryAPIRequest {
	return &TaobaoTopEventSubscriptionQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopEventSubscriptionQueryAPIRequest) GetApiMethodName() string {
	return "taobao.top.event.subscription.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopEventSubscriptionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopEventSubscriptionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTriggerCode is TriggerCode Setter
// 连接平台触发事件TriggerCode
func (r *TaobaoTopEventSubscriptionQueryAPIRequest) SetTriggerCode(_triggerCode string) error {
	r._triggerCode = _triggerCode
	r.Set("trigger_code", _triggerCode)
	return nil
}

// GetTriggerCode TriggerCode Getter
func (r TaobaoTopEventSubscriptionQueryAPIRequest) GetTriggerCode() string {
	return r._triggerCode
}
