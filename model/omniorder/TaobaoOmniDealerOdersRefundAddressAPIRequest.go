package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomnidealerodersrefundaddressAPIRequest 经销商查询逆向退货地址 API请求
// taobao.omni.dealer.oders.refund.address
//
// 经销商查询逆向退货地址
type TaobaoomnidealerodersrefundaddressAPIRequest struct {
	model.Params
	// 经销商子订单id
	_orderId int64
}

// NewTaobaoomnidealerodersrefundaddressRequest 初始化TaobaoomnidealerodersrefundaddressAPIRequest对象
func NewTaobaoomnidealerodersrefundaddressRequest() *TaobaoomnidealerodersrefundaddressAPIRequest {
	return &TaobaoomnidealerodersrefundaddressAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomnidealerodersrefundaddressAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.refund.address"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomnidealerodersrefundaddressAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomnidealerodersrefundaddressAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 经销商子订单id
func (r *TaobaoomnidealerodersrefundaddressAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoomnidealerodersrefundaddressAPIRequest) GetOrderId() int64 {
	return r._orderId
}
