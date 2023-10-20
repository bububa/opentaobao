package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripiebuyerorderbookpayAPIRequest 【国际机票】下单预定支付 API请求
// alitrip.ie.buyer.order.bookpay
//
// 【国际机票】 生单预定支付接口
type AlitripiebuyerorderbookpayAPIRequest struct {
	model.Params
	// 生单支付请求参数
	_bookPayOrderParam *BookPayOrderRq
}

// NewAlitripiebuyerorderbookpayRequest 初始化AlitripiebuyerorderbookpayAPIRequest对象
func NewAlitripiebuyerorderbookpayRequest() *AlitripiebuyerorderbookpayAPIRequest {
	return &AlitripiebuyerorderbookpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripiebuyerorderbookpayAPIRequest) GetApiMethodName() string {
	return "alitrip.ie.buyer.order.bookpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripiebuyerorderbookpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripiebuyerorderbookpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBookPayOrderParam is BookPayOrderParam Setter
// 生单支付请求参数
func (r *AlitripiebuyerorderbookpayAPIRequest) SetBookPayOrderParam(_bookPayOrderParam *BookPayOrderRq) error {
	r._bookPayOrderParam = _bookPayOrderParam
	r.Set("book_pay_order_param", _bookPayOrderParam)
	return nil
}

// GetBookPayOrderParam BookPayOrderParam Getter
func (r AlitripiebuyerorderbookpayAPIRequest) GetBookPayOrderParam() *BookPayOrderRq {
	return r._bookPayOrderParam
}
