package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayPartnerOrderQueryAPIRequest 商户查询订单 API请求
// taobao.tvpay.partner.order.query
//
// 给商户提供的查询订单状态的API
type TaobaoTvpayPartnerOrderQueryAPIRequest struct {
	model.Params
	// 商户订单号
	_orderNo string
}

// NewTaobaoTvpayPartnerOrderQueryRequest 初始化TaobaoTvpayPartnerOrderQueryAPIRequest对象
func NewTaobaoTvpayPartnerOrderQueryRequest() *TaobaoTvpayPartnerOrderQueryAPIRequest {
	return &TaobaoTvpayPartnerOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayPartnerOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.partner.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayPartnerOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderNo is OrderNo Setter
// 商户订单号
func (r *TaobaoTvpayPartnerOrderQueryAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaoTvpayPartnerOrderQueryAPIRequest) GetOrderNo() string {
	return r._orderNo
}
