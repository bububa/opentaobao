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
type AliexpressTradeOrderOpenCheckAPIRequest struct {
    model.Params
    // 预下单入参
    _paramPreCreateOrderRequest   *PreCreateOrderRequest
    // 客户端信息
    _paramClientInfo   *ClientInfo
}

// 初始化AliexpressTradeOrderOpenCheckAPIRequest对象
func NewAliexpressTradeOrderOpenCheckRequest() *AliexpressTradeOrderOpenCheckAPIRequest{
    return &AliexpressTradeOrderOpenCheckAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetApiMethodName() string {
    return "aliexpress.trade.order.open.check"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPreCreateOrderRequest Setter
// 预下单入参
func (r *AliexpressTradeOrderOpenCheckAPIRequest) SetParamPreCreateOrderRequest(_paramPreCreateOrderRequest *PreCreateOrderRequest) error {
    r._paramPreCreateOrderRequest = _paramPreCreateOrderRequest
    r.Set("param_pre_create_order_request", _paramPreCreateOrderRequest)
    return nil
}

// ParamPreCreateOrderRequest Getter
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetParamPreCreateOrderRequest() *PreCreateOrderRequest {
    return r._paramPreCreateOrderRequest
}
// ParamClientInfo Setter
// 客户端信息
func (r *AliexpressTradeOrderOpenCheckAPIRequest) SetParamClientInfo(_paramClientInfo *ClientInfo) error {
    r._paramClientInfo = _paramClientInfo
    r.Set("param_client_info", _paramClientInfo)
    return nil
}

// ParamClientInfo Getter
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetParamClientInfo() *ClientInfo {
    return r._paramClientInfo
}
