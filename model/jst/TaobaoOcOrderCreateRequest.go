package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OC订单 APIRequest
taobao.oc.order.create

创建OC订单接口
*/
type TaobaoOcOrderCreateRequest struct {
    model.Params

    // OC订单
    paramOCOrder   *OcOrder 

}

func NewTaobaoOcOrderCreateRequest() *TaobaoOcOrderCreateRequest{
    return &TaobaoOcOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOcOrderCreateRequest) GetApiMethodName() string {
    return "taobao.oc.order.create"
}

func (r TaobaoOcOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOcOrderCreateRequest) SetParamOCOrder(paramOCOrder *OcOrder) error {
    r.paramOCOrder = paramOCOrder
    r.Set("param_o_c_order", paramOCOrder)
    return nil
}

func (r TaobaoOcOrderCreateRequest) GetParamOCOrder() *OcOrder {
    return r.paramOCOrder
}

