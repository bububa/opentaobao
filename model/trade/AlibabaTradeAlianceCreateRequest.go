package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推客平台订单回流 APIRequest
alibaba.trade.aliance.create

推客平台订单回流
*/
type AlibabaTradeAlianceCreateRequest struct {
    model.Params

    // 下单请求
    paramIsvCreateOrderParam   *IsvCreateOrderParam 

}

func NewAlibabaTradeAlianceCreateRequest() *AlibabaTradeAlianceCreateRequest{
    return &AlibabaTradeAlianceCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTradeAlianceCreateRequest) GetApiMethodName() string {
    return "alibaba.trade.aliance.create"
}

func (r AlibabaTradeAlianceCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTradeAlianceCreateRequest) SetParamIsvCreateOrderParam(paramIsvCreateOrderParam *IsvCreateOrderParam) error {
    r.paramIsvCreateOrderParam = paramIsvCreateOrderParam
    r.Set("param_isv_create_order_param", paramIsvCreateOrderParam)
    return nil
}

func (r AlibabaTradeAlianceCreateRequest) GetParamIsvCreateOrderParam() *IsvCreateOrderParam {
    return r.paramIsvCreateOrderParam
}

