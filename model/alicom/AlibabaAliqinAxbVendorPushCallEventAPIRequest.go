package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinaxbvendorpushcalleventAPIRequest 呼叫事件推送 API请求
// alibaba.aliqin.axb.vendor.push.call.event
//
// 呼叫事件推送
// 响铃时间、摘机事件
type AlibabaaliqinaxbvendorpushcalleventAPIRequest struct {
	model.Params
	// 呼叫事件推送请求
	_eventCallRequest *EventCallRequest
}

// NewAlibabaaliqinaxbvendorpushcalleventRequest 初始化AlibabaaliqinaxbvendorpushcalleventAPIRequest对象
func NewAlibabaaliqinaxbvendorpushcalleventRequest() *AlibabaaliqinaxbvendorpushcalleventAPIRequest {
	return &AlibabaaliqinaxbvendorpushcalleventAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinaxbvendorpushcalleventAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.push.call.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinaxbvendorpushcalleventAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinaxbvendorpushcalleventAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEventCallRequest is EventCallRequest Setter
// 呼叫事件推送请求
func (r *AlibabaaliqinaxbvendorpushcalleventAPIRequest) SetEventCallRequest(_eventCallRequest *EventCallRequest) error {
	r._eventCallRequest = _eventCallRequest
	r.Set("event_call_request", _eventCallRequest)
	return nil
}

// GetEventCallRequest EventCallRequest Getter
func (r AlibabaaliqinaxbvendorpushcalleventAPIRequest) GetEventCallRequest() *EventCallRequest {
	return r._eventCallRequest
}
