package ieagency

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripIeBuyerOrderBookpayAPIRequest) Reset() {
	r._bookPayOrderParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripIeBuyerOrderBookpayAPIRequest) GetApiMethodName() string {
	return "alitrip.ie.buyer.order.bookpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripIeBuyerOrderBookpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripIeBuyerOrderBookpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBookPayOrderParam is BookPayOrderParam Setter
// 生单支付请求参数
func (r *AlitripIeBuyerOrderBookpayAPIRequest) SetBookPayOrderParam(_bookPayOrderParam *BookPayOrderRq) error {
	r._bookPayOrderParam = _bookPayOrderParam
	r.Set("book_pay_order_param", _bookPayOrderParam)
	return nil
}

// GetBookPayOrderParam BookPayOrderParam Getter
func (r AlitripIeBuyerOrderBookpayAPIRequest) GetBookPayOrderParam() *BookPayOrderRq {
	return r._bookPayOrderParam
}

var poolAlitripIeBuyerOrderBookpayAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripIeBuyerOrderBookpayRequest()
	},
}

// GetAlitripIeBuyerOrderBookpayRequest 从 sync.Pool 获取 AlitripIeBuyerOrderBookpayAPIRequest
func GetAlitripIeBuyerOrderBookpayAPIRequest() *AlitripIeBuyerOrderBookpayAPIRequest {
	return poolAlitripIeBuyerOrderBookpayAPIRequest.Get().(*AlitripIeBuyerOrderBookpayAPIRequest)
}

// ReleaseAlitripIeBuyerOrderBookpayAPIRequest 将 AlitripIeBuyerOrderBookpayAPIRequest 放入 sync.Pool
func ReleaseAlitripIeBuyerOrderBookpayAPIRequest(v *AlitripIeBuyerOrderBookpayAPIRequest) {
	v.Reset()
	poolAlitripIeBuyerOrderBookpayAPIRequest.Put(v)
}
