package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】下单预定支付 API请求
alitrip.ie.buyer.order.bookpay

【国际机票】 生单预定支付接口
*/
type AlitripIeBuyerOrderBookpayRequest struct {
    model.Params
    // 生单支付请求参数
    _bookPayOrderParam   *BookPayOrderRq
}

// 初始化AlitripIeBuyerOrderBookpayRequest对象
func NewAlitripIeBuyerOrderBookpayRequest() *AlitripIeBuyerOrderBookpayRequest{
    return &AlitripIeBuyerOrderBookpayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripIeBuyerOrderBookpayRequest) GetApiMethodName() string {
    return "alitrip.ie.buyer.order.bookpay"
}

// IRequest interface 方法, 获取API参数
func (r AlitripIeBuyerOrderBookpayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BookPayOrderParam Setter
// 生单支付请求参数
func (r *AlitripIeBuyerOrderBookpayRequest) SetBookPayOrderParam(_bookPayOrderParam *BookPayOrderRq) error {
    r._bookPayOrderParam = _bookPayOrderParam
    r.Set("book_pay_order_param", _bookPayOrderParam)
    return nil
}

// BookPayOrderParam Getter
func (r AlitripIeBuyerOrderBookpayRequest) GetBookPayOrderParam() *BookPayOrderRq {
    return r._bookPayOrderParam
}
