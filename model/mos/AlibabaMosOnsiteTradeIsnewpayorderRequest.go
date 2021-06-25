package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
是否为新支付订单 APIRequest
alibaba.mos.onsite.trade.isnewpayorder

退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
*/
type AlibabaMosOnsiteTradeIsnewpayorderRequest struct {
    model.Params

    // 外部订单号
    outTradeNo   string 

}

func NewAlibabaMosOnsiteTradeIsnewpayorderRequest() *AlibabaMosOnsiteTradeIsnewpayorderRequest{
    return &AlibabaMosOnsiteTradeIsnewpayorderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOnsiteTradeIsnewpayorderRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.isnewpayorder"
}

func (r AlibabaMosOnsiteTradeIsnewpayorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOnsiteTradeIsnewpayorderRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

func (r AlibabaMosOnsiteTradeIsnewpayorderRequest) GetOutTradeNo() string {
    return r.outTradeNo
}

