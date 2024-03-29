package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderStatusAPIRequest 商家订单状态改变通知接口（神州专车接口） API请求
// taobao.alitrip.car.order.status
//
// 商家订单状态改变通知接口，神州专车专用接口！
type TaobaoAlitripCarOrderStatusAPIRequest struct {
	model.Params
	// 固定值：statusChanged
	_operation string
	// 飞猪订单ID
	_orderId string
	// 服务商ID
	_providerId string
	// 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
	_status string
}

// NewTaobaoAlitripCarOrderStatusRequest 初始化TaobaoAlitripCarOrderStatusAPIRequest对象
func NewTaobaoAlitripCarOrderStatusRequest() *TaobaoAlitripCarOrderStatusAPIRequest {
	return &TaobaoAlitripCarOrderStatusAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarOrderStatusAPIRequest) Reset() {
	r._operation = ""
	r._orderId = ""
	r._providerId = ""
	r._status = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderStatusAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarOrderStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperation is Operation Setter
// 固定值：statusChanged
func (r *TaobaoAlitripCarOrderStatusAPIRequest) SetOperation(_operation string) error {
	r._operation = _operation
	r.Set("operation", _operation)
	return nil
}

// GetOperation Operation Getter
func (r TaobaoAlitripCarOrderStatusAPIRequest) GetOperation() string {
	return r._operation
}

// SetOrderId is OrderId Setter
// 飞猪订单ID
func (r *TaobaoAlitripCarOrderStatusAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripCarOrderStatusAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetProviderId is ProviderId Setter
// 服务商ID
func (r *TaobaoAlitripCarOrderStatusAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r TaobaoAlitripCarOrderStatusAPIRequest) GetProviderId() string {
	return r._providerId
}

// SetStatus is Status Setter
// 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
func (r *TaobaoAlitripCarOrderStatusAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoAlitripCarOrderStatusAPIRequest) GetStatus() string {
	return r._status
}

var poolTaobaoAlitripCarOrderStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarOrderStatusRequest()
	},
}

// GetTaobaoAlitripCarOrderStatusRequest 从 sync.Pool 获取 TaobaoAlitripCarOrderStatusAPIRequest
func GetTaobaoAlitripCarOrderStatusAPIRequest() *TaobaoAlitripCarOrderStatusAPIRequest {
	return poolTaobaoAlitripCarOrderStatusAPIRequest.Get().(*TaobaoAlitripCarOrderStatusAPIRequest)
}

// ReleaseTaobaoAlitripCarOrderStatusAPIRequest 将 TaobaoAlitripCarOrderStatusAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarOrderStatusAPIRequest(v *TaobaoAlitripCarOrderStatusAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarOrderStatusAPIRequest.Put(v)
}
