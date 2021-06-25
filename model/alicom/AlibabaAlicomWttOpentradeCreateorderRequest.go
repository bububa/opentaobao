package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
充值送活动下单接口 APIRequest
alibaba.alicom.wtt.opentrade.createorder

提供给话费宝创建淘宝订单
*/
type AlibabaAlicomWttOpentradeCreateorderRequest struct {
    model.Params

    // 入参请求说明
    param0   *OpentradCreateOrderRequestDTO 

}

func NewAlibabaAlicomWttOpentradeCreateorderRequest() *AlibabaAlicomWttOpentradeCreateorderRequest{
    return &AlibabaAlicomWttOpentradeCreateorderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlicomWttOpentradeCreateorderRequest) GetApiMethodName() string {
    return "alibaba.alicom.wtt.opentrade.createorder"
}

func (r AlibabaAlicomWttOpentradeCreateorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlicomWttOpentradeCreateorderRequest) SetParam0(param0 *OpentradCreateOrderRequestDTO) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaAlicomWttOpentradeCreateorderRequest) GetParam0() *OpentradCreateOrderRequestDTO {
    return r.param0
}

