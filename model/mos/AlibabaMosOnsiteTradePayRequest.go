package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新商场当面付商户扫码付 APIRequest
alibaba.mos.onsite.trade.pay

收银员使用扫码设备读取用户“付款码”后，将二维码或条码信息通过本接口上送至喵街发起支付。
*/
type AlibabaMosOnsiteTradePayRequest struct {
    model.Params

    // 创建订单请求参数
    onsiteTradePayRequest   *OnsiteTradePayRequest 

}

func NewAlibabaMosOnsiteTradePayRequest() *AlibabaMosOnsiteTradePayRequest{
    return &AlibabaMosOnsiteTradePayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOnsiteTradePayRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.pay"
}

func (r AlibabaMosOnsiteTradePayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOnsiteTradePayRequest) SetOnsiteTradePayRequest(onsiteTradePayRequest *OnsiteTradePayRequest) error {
    r.onsiteTradePayRequest = onsiteTradePayRequest
    r.Set("onsite_trade_pay_request", onsiteTradePayRequest)
    return nil
}

func (r AlibabaMosOnsiteTradePayRequest) GetOnsiteTradePayRequest() *OnsiteTradePayRequest {
    return r.onsiteTradePayRequest
}

