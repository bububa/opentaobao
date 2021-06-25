package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户查询订单 APIRequest
taobao.tvpay.partner.order.query

给商户提供的查询订单状态的API
*/
type TaobaoTvpayPartnerOrderQueryRequest struct {
    model.Params

    // 商户订单号
    orderNo   string 

}

func NewTaobaoTvpayPartnerOrderQueryRequest() *TaobaoTvpayPartnerOrderQueryRequest{
    return &TaobaoTvpayPartnerOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayPartnerOrderQueryRequest) GetApiMethodName() string {
    return "taobao.tvpay.partner.order.query"
}

func (r TaobaoTvpayPartnerOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayPartnerOrderQueryRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r TaobaoTvpayPartnerOrderQueryRequest) GetOrderNo() string {
    return r.orderNo
}

