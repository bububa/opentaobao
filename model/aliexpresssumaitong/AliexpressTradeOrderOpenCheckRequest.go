package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Aliexpress开放平台下单前置检查 APIRequest
aliexpress.trade.order.open.check

Aliexpress开放平台下单前通过下单入参获取token
*/
type AliexpressTradeOrderOpenCheckRequest struct {
    model.Params

    // 预下单入参
    paramPreCreateOrderRequest   *PreCreateOrderRequest 

    // 客户端信息
    paramClientInfo   *ClientInfo 

}

func NewAliexpressTradeOrderOpenCheckRequest() *AliexpressTradeOrderOpenCheckRequest{
    return &AliexpressTradeOrderOpenCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressTradeOrderOpenCheckRequest) GetApiMethodName() string {
    return "aliexpress.trade.order.open.check"
}

func (r AliexpressTradeOrderOpenCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressTradeOrderOpenCheckRequest) SetParamPreCreateOrderRequest(paramPreCreateOrderRequest *PreCreateOrderRequest) error {
    r.paramPreCreateOrderRequest = paramPreCreateOrderRequest
    r.Set("param_pre_create_order_request", paramPreCreateOrderRequest)
    return nil
}

func (r AliexpressTradeOrderOpenCheckRequest) GetParamPreCreateOrderRequest() *PreCreateOrderRequest {
    return r.paramPreCreateOrderRequest
}

func (r *AliexpressTradeOrderOpenCheckRequest) SetParamClientInfo(paramClientInfo *ClientInfo) error {
    r.paramClientInfo = paramClientInfo
    r.Set("param_client_info", paramClientInfo)
    return nil
}

func (r AliexpressTradeOrderOpenCheckRequest) GetParamClientInfo() *ClientInfo {
    return r.paramClientInfo
}

