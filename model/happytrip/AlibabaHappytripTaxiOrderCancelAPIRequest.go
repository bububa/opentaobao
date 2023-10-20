package happytrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiOrderCancelAPIRequest) Reset() {
	r._orderId = ""
	r._force = ""
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaHappytripTaxiOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiOrderCancelRequest()
	},
}

// GetAlibabaHappytripTaxiOrderCancelRequest 从 sync.Pool 获取 AlibabaHappytripTaxiOrderCancelAPIRequest
func GetAlibabaHappytripTaxiOrderCancelAPIRequest() *AlibabaHappytripTaxiOrderCancelAPIRequest {
	return poolAlibabaHappytripTaxiOrderCancelAPIRequest.Get().(*AlibabaHappytripTaxiOrderCancelAPIRequest)
}

// ReleaseAlibabaHappytripTaxiOrderCancelAPIRequest 将 AlibabaHappytripTaxiOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderCancelAPIRequest(v *AlibabaHappytripTaxiOrderCancelAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderCancelAPIRequest.Put(v)
}
