package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】下单预定支付 APIRequest
alitrip.ie.buyer.order.bookpay

【国际机票】 生单预定支付接口
*/
type AlitripIeBuyerOrderBookpayRequest struct {
    model.Params

    // 生单支付请求参数
    bookPayOrderParam   *BookPayOrderRq 

}

func NewAlitripIeBuyerOrderBookpayRequest() *AlitripIeBuyerOrderBookpayRequest{
    return &AlitripIeBuyerOrderBookpayRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripIeBuyerOrderBookpayRequest) GetApiMethodName() string {
    return "alitrip.ie.buyer.order.bookpay"
}

func (r AlitripIeBuyerOrderBookpayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripIeBuyerOrderBookpayRequest) SetBookPayOrderParam(bookPayOrderParam *BookPayOrderRq) error {
    r.bookPayOrderParam = bookPayOrderParam
    r.Set("book_pay_order_param", bookPayOrderParam)
    return nil
}

func (r AlitripIeBuyerOrderBookpayRequest) GetBookPayOrderParam() *BookPayOrderRq {
    return r.bookPayOrderParam
}

