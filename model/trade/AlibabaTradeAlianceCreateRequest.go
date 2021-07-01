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
type AlibabaTradeAlianceCreateAPIRequest struct {
    model.Params
    // 下单请求
    _paramIsvCreateOrderParam   *IsvCreateOrderParam
}

// 初始化AlibabaTradeAlianceCreateAPIRequest对象
func NewAlibabaTradeAlianceCreateRequest() *AlibabaTradeAlianceCreateAPIRequest{
    return &AlibabaTradeAlianceCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTradeAlianceCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.trade.aliance.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTradeAlianceCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamIsvCreateOrderParam Setter
// 下单请求
func (r *AlibabaTradeAlianceCreateAPIRequest) SetParamIsvCreateOrderParam(_paramIsvCreateOrderParam *IsvCreateOrderParam) error {
    r._paramIsvCreateOrderParam = _paramIsvCreateOrderParam
    r.Set("param_isv_create_order_param", _paramIsvCreateOrderParam)
    return nil
}

// ParamIsvCreateOrderParam Getter
func (r AlibabaTradeAlianceCreateAPIRequest) GetParamIsvCreateOrderParam() *IsvCreateOrderParam {
    return r._paramIsvCreateOrderParam
}
