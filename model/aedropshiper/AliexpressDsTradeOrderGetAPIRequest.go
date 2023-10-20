package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdstradeordergetAPIRequest 交易订单查询 API请求
// aliexpress.ds.trade.order.get
//
// 交易订单查询
type AliexpressdstradeordergetAPIRequest struct {
	model.Params
	// AE order id
	_orderId int64
}

// NewAliexpressdstradeordergetRequest 初始化AliexpressdstradeordergetAPIRequest对象
func NewAliexpressdstradeordergetRequest() *AliexpressdstradeordergetAPIRequest {
	return &AliexpressdstradeordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressdstradeordergetAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.trade.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressdstradeordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressdstradeordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// AE order id
func (r *AliexpressdstradeordergetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AliexpressdstradeordergetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
