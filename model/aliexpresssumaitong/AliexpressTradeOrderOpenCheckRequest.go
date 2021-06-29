package aliexpresssumaitong

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Aliexpress开放平台下单前置检查 API请求
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

// 初始化AliexpressTradeOrderOpenCheckRequest对象
func NewAliexpressTradeOrderOpenCheckRequest() *AliexpressTradeOrderOpenCheckRequest{
    return &AliexpressTradeOrderOpenCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTradeOrderOpenCheckRequest) GetApiMethodName() string {
    return "aliexpress.trade.order.open.check"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTradeOrderOpenCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPreCreateOrderRequest Setter
// 预下单入参
func (r *AliexpressTradeOrderOpenCheckRequest) SetParamPreCreateOrderRequest(paramPreCreateOrderRequest *PreCreateOrderRequest) error {
    r.paramPreCreateOrderRequest = paramPreCreateOrderRequest
    r.Set("param_pre_create_order_request", paramPreCreateOrderRequest)
    return nil
}

// ParamPreCreateOrderRequest Getter
func (r AliexpressTradeOrderOpenCheckRequest) GetParamPreCreateOrderRequest() *PreCreateOrderRequest {
    return r.paramPreCreateOrderRequest
}
// ParamClientInfo Setter
// 客户端信息
func (r *AliexpressTradeOrderOpenCheckRequest) SetParamClientInfo(paramClientInfo *ClientInfo) error {
    r.paramClientInfo = paramClientInfo
    r.Set("param_client_info", paramClientInfo)
    return nil
}

// ParamClientInfo Getter
func (r AliexpressTradeOrderOpenCheckRequest) GetParamClientInfo() *ClientInfo {
    return r.paramClientInfo
}
