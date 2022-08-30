package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorPushCallEventAPIRequest 呼叫事件推送 API请求
// alibaba.aliqin.axb.vendor.push.call.event
//
// 呼叫事件推送
// 响铃时间、摘机事件
type AlibabaAliqinAxbVendorPushCallEventAPIRequest struct {
	model.Params
	// 呼叫事件推送请求
	_eventCallRequest *EventCallRequest
}

// NewAlibabaAliqinAxbVendorPushCallEventRequest 初始化AlibabaAliqinAxbVendorPushCallEventAPIRequest对象
func NewAlibabaAliqinAxbVendorPushCallEventRequest() *AlibabaAliqinAxbVendorPushCallEventAPIRequest {
	return &AlibabaAliqinAxbVendorPushCallEventAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorPushCallEventAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.push.call.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorPushCallEventAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEventCallRequest is EventCallRequest Setter
// 呼叫事件推送请求
func (r *AlibabaAliqinAxbVendorPushCallEventAPIRequest) SetEventCallRequest(_eventCallRequest *EventCallRequest) error {
	r._eventCallRequest = _eventCallRequest
	r.Set("event_call_request", _eventCallRequest)
	return nil
}

// GetEventCallRequest EventCallRequest Getter
func (r AlibabaAliqinAxbVendorPushCallEventAPIRequest) GetEventCallRequest() *EventCallRequest {
	return r._eventCallRequest
}
