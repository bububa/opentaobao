package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaorefundrefundactionsgetAPIRequest 判断该订单能执行的逆向操作集合列表 API请求
// cainiao.refund.refundactions.get
//
// 判断该订单能执行的逆向操作集合列表
type CainiaorefundrefundactionsgetAPIRequest struct {
	model.Params
	// 子订单ID
	_orderId string
}

// NewCainiaorefundrefundactionsgetRequest 初始化CainiaorefundrefundactionsgetAPIRequest对象
func NewCainiaorefundrefundactionsgetRequest() *CainiaorefundrefundactionsgetAPIRequest {
	return &CainiaorefundrefundactionsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaorefundrefundactionsgetAPIRequest) GetApiMethodName() string {
	return "cainiao.refund.refundactions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaorefundrefundactionsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaorefundrefundactionsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 子订单ID
func (r *CainiaorefundrefundactionsgetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r CainiaorefundrefundactionsgetAPIRequest) GetOrderId() string {
	return r._orderId
}
