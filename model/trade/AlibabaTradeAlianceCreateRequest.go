package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推客平台订单回流 API请求
alibaba.trade.aliance.create

推客平台订单回流
*/
type AlibabaTradeAlianceCreateRequest struct {
    model.Params
    // 下单请求
    paramIsvCreateOrderParam   *IsvCreateOrderParam
}

// 初始化AlibabaTradeAlianceCreateRequest对象
func NewAlibabaTradeAlianceCreateRequest() *AlibabaTradeAlianceCreateRequest{
    return &AlibabaTradeAlianceCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTradeAlianceCreateRequest) GetApiMethodName() string {
    return "alibaba.trade.aliance.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTradeAlianceCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamIsvCreateOrderParam Setter
// 下单请求
func (r *AlibabaTradeAlianceCreateRequest) SetParamIsvCreateOrderParam(paramIsvCreateOrderParam *IsvCreateOrderParam) error {
    r.paramIsvCreateOrderParam = paramIsvCreateOrderParam
    r.Set("param_isv_create_order_param", paramIsvCreateOrderParam)
    return nil
}

// ParamIsvCreateOrderParam Getter
func (r AlibabaTradeAlianceCreateRequest) GetParamIsvCreateOrderParam() *IsvCreateOrderParam {
    return r.paramIsvCreateOrderParam
}
