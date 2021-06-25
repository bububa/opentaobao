package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下新退款接口（专为老退款接口调用） APIRequest
alibaba.mos.onsite.trade.oldrefund

线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
*/
type AlibabaMosOnsiteTradeOldrefundRequest struct {
    model.Params

    // 交易退款请求
    onsiteRefundRequest   *OnsiteRefundRequest 

}

func NewAlibabaMosOnsiteTradeOldrefundRequest() *AlibabaMosOnsiteTradeOldrefundRequest{
    return &AlibabaMosOnsiteTradeOldrefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOnsiteTradeOldrefundRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.oldrefund"
}

func (r AlibabaMosOnsiteTradeOldrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOnsiteTradeOldrefundRequest) SetOnsiteRefundRequest(onsiteRefundRequest *OnsiteRefundRequest) error {
    r.onsiteRefundRequest = onsiteRefundRequest
    r.Set("onsite_refund_request", onsiteRefundRequest)
    return nil
}

func (r AlibabaMosOnsiteTradeOldrefundRequest) GetOnsiteRefundRequest() *OnsiteRefundRequest {
    return r.onsiteRefundRequest
}

