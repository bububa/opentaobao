package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiordercancelAPIRequest 取消叫车 API请求
// alibaba.happytrip.taxi.order.cancel
//
// 取消叫车订单,行程中的订单不能取消
type AlibabahappytriptaxiordercancelAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 是否强制取消(true 或 false)默认false
	_force string
	// 取消类型，0：系统取消，非0：用户取消
	_type int64
}

// NewAlibabahappytriptaxiordercancelRequest 初始化AlibabahappytriptaxiordercancelAPIRequest对象
func NewAlibabahappytriptaxiordercancelRequest() *AlibabahappytriptaxiordercancelAPIRequest {
	return &AlibabahappytriptaxiordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahappytriptaxiordercancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiordercancelAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetForce is Force Setter
// 是否强制取消(true 或 false)默认false
func (r *AlibabahappytriptaxiordercancelAPIRequest) SetForce(_force string) error {
	r._force = _force
	r.Set("force", _force)
	return nil
}

// GetForce Force Getter
func (r AlibabahappytriptaxiordercancelAPIRequest) GetForce() string {
	return r._force
}

// SetType is Type Setter
// 取消类型，0：系统取消，非0：用户取消
func (r *AlibabahappytriptaxiordercancelAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabahappytriptaxiordercancelAPIRequest) GetType() int64 {
	return r._type
}
