package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseordergetAPIRequest 获取租赁订单信息 API请求
// tmall.car.leaseorder.get
//
// 卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家
type TmallcarleaseordergetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// NewTmallcarleaseordergetRequest 初始化TmallcarleaseordergetAPIRequest对象
func NewTmallcarleaseordergetRequest() *TmallcarleaseordergetAPIRequest {
	return &TmallcarleaseordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleaseordergetAPIRequest) GetApiMethodName() string {
	return "tmall.car.leaseorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleaseordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleaseordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TmallcarleaseordergetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallcarleaseordergetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
