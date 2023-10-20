package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderQueryAPIRequest 查询订单 API请求
// alibaba.idle.rent.order.query
//
// 查询订单信息
type AlibabaIdleRentOrderQueryAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabaIdleRentOrderQueryRequest 初始化AlibabaIdleRentOrderQueryAPIRequest对象
func NewAlibabaIdleRentOrderQueryRequest() *AlibabaIdleRentOrderQueryAPIRequest {
	return &AlibabaIdleRentOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderQueryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaIdleRentOrderQueryAPIRequest) GetOrderId() int64 {
	return r._orderId
}
