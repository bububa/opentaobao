package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按OC订单分账 APIRequest
taobao.oc.order.ap.update

对OC订单执行分账操作
*/
type TaobaoOcOrderApUpdateRequest struct {
    model.Params

    // 调用创建OC订单接口生成的id
    ocOrderId   int64 

}

func NewTaobaoOcOrderApUpdateRequest() *TaobaoOcOrderApUpdateRequest{
    return &TaobaoOcOrderApUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOcOrderApUpdateRequest) GetApiMethodName() string {
    return "taobao.oc.order.ap.update"
}

func (r TaobaoOcOrderApUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOcOrderApUpdateRequest) SetOcOrderId(ocOrderId int64) error {
    r.ocOrderId = ocOrderId
    r.Set("oc_order_id", ocOrderId)
    return nil
}

func (r TaobaoOcOrderApUpdateRequest) GetOcOrderId() int64 {
    return r.ocOrderId
}

