package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinAxbVendorPushCallEventAPIRequest) Reset() {
	r._eventCallRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorPushCallEventAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.axb.vendor.push.call.event"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorPushCallEventAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinAxbVendorPushCallEventAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAliqinAxbVendorPushCallEventAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinAxbVendorPushCallEventRequest()
	},
}

// GetAlibabaAliqinAxbVendorPushCallEventRequest 从 sync.Pool 获取 AlibabaAliqinAxbVendorPushCallEventAPIRequest
func GetAlibabaAliqinAxbVendorPushCallEventAPIRequest() *AlibabaAliqinAxbVendorPushCallEventAPIRequest {
	return poolAlibabaAliqinAxbVendorPushCallEventAPIRequest.Get().(*AlibabaAliqinAxbVendorPushCallEventAPIRequest)
}

// ReleaseAlibabaAliqinAxbVendorPushCallEventAPIRequest 将 AlibabaAliqinAxbVendorPushCallEventAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinAxbVendorPushCallEventAPIRequest(v *AlibabaAliqinAxbVendorPushCallEventAPIRequest) {
	v.Reset()
	poolAlibabaAliqinAxbVendorPushCallEventAPIRequest.Put(v)
}
