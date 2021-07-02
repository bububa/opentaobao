package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripIeBuyerOrderBookpayAPIRequest 【国际机票】下单预定支付 API请求
// alitrip.ie.buyer.order.bookpay
//
// 【国际机票】 生单预定支付接口
type AlitripIeBuyerOrderBookpayAPIRequest struct {
	model.Params
	// 生单支付请求参数
	_bookPayOrderParam *BookPayOrderRq
}

// NewAlitripIeBuyerOrderBookpayRequest 初始化AlitripIeBuyerOrderBookpayAPIRequest对象
func NewAlitripIeBuyerOrderBookpayRequest() *AlitripIeBuyerOrderBookpayAPIRequest {
	return &AlitripIeBuyerOrderBookpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripIeBuyerOrderBookpayAPIRequest) GetApiMethodName() string {
	return "alitrip.ie.buyer.order.bookpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripIeBuyerOrderBookpayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BookPayOrderParam Setter
// 生单支付请求参数
func (r *AlitripIeBuyerOrderBookpayAPIRequest) SetBookPayOrderParam(_bookPayOrderParam *BookPayOrderRq) error {
	r._bookPayOrderParam = _bookPayOrderParam
	r.Set("book_pay_order_param", _bookPayOrderParam)
	return nil
}

// Get BookPayOrderParam Getter
func (r AlitripIeBuyerOrderBookpayAPIRequest) GetBookPayOrderParam() *BookPayOrderRq {
	return r._bookPayOrderParam
}
