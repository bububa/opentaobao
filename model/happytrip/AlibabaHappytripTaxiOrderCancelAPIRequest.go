package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderCancelAPIRequest 取消叫车 API请求
// alibaba.happytrip.taxi.order.cancel
//
// 取消叫车订单,行程中的订单不能取消
type AlibabaHappytripTaxiOrderCancelAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 是否强制取消(true 或 false)默认false
	_force string
	// 取消类型，0：系统取消，非0：用户取消
	_type int64
}

// NewAlibabaHappytripTaxiOrderCancelRequest 初始化AlibabaHappytripTaxiOrderCancelAPIRequest对象
func NewAlibabaHappytripTaxiOrderCancelRequest() *AlibabaHappytripTaxiOrderCancelAPIRequest {
	return &AlibabaHappytripTaxiOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaHappytripTaxiOrderCancelAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetForce is Force Setter
// 是否强制取消(true 或 false)默认false
func (r *AlibabaHappytripTaxiOrderCancelAPIRequest) SetForce(_force string) error {
	r._force = _force
	r.Set("force", _force)
	return nil
}

// GetForce Force Getter
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetForce() string {
	return r._force
}

// SetType is Type Setter
// 取消类型，0：系统取消，非0：用户取消
func (r *AlibabaHappytripTaxiOrderCancelAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetType() int64 {
	return r._type
}
