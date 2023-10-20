package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarorderstatusAPIRequest 商家订单状态改变通知接口（神州专车接口） API请求
// taobao.alitrip.car.order.status
//
// 商家订单状态改变通知接口，神州专车专用接口！
type TaobaoalitripcarorderstatusAPIRequest struct {
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

// NewTaobaoalitripcarorderstatusRequest 初始化TaobaoalitripcarorderstatusAPIRequest对象
func NewTaobaoalitripcarorderstatusRequest() *TaobaoalitripcarorderstatusAPIRequest {
	return &TaobaoalitripcarorderstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcarorderstatusAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcarorderstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcarorderstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperation is Operation Setter
// 固定值：statusChanged
func (r *TaobaoalitripcarorderstatusAPIRequest) SetOperation(_operation string) error {
	r._operation = _operation
	r.Set("operation", _operation)
	return nil
}

// GetOperation Operation Getter
func (r TaobaoalitripcarorderstatusAPIRequest) GetOperation() string {
	return r._operation
}

// SetOrderId is OrderId Setter
// 飞猪订单ID
func (r *TaobaoalitripcarorderstatusAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoalitripcarorderstatusAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetProviderId is ProviderId Setter
// 服务商ID
func (r *TaobaoalitripcarorderstatusAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r TaobaoalitripcarorderstatusAPIRequest) GetProviderId() string {
	return r._providerId
}

// SetStatus is Status Setter
// 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
func (r *TaobaoalitripcarorderstatusAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoalitripcarorderstatusAPIRequest) GetStatus() string {
	return r._status
}
