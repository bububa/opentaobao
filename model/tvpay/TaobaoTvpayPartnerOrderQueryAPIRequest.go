package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotvpaypartnerorderqueryAPIRequest 商户查询订单 API请求
// taobao.tvpay.partner.order.query
//
// 给商户提供的查询订单状态的API
type TaobaotvpaypartnerorderqueryAPIRequest struct {
	model.Params
	// 商户订单号
	_orderNo string
}

// NewTaobaotvpaypartnerorderqueryRequest 初始化TaobaotvpaypartnerorderqueryAPIRequest对象
func NewTaobaotvpaypartnerorderqueryRequest() *TaobaotvpaypartnerorderqueryAPIRequest {
	return &TaobaotvpaypartnerorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotvpaypartnerorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.partner.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotvpaypartnerorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotvpaypartnerorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 商户订单号
func (r *TaobaotvpaypartnerorderqueryAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaotvpaypartnerorderqueryAPIRequest) GetOrderNo() string {
	return r._orderNo
}
