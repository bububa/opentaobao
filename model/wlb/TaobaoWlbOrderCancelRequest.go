package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消物流宝订单 APIRequest
taobao.wlb.order.cancel

取消物流宝订单
*/
type TaobaoWlbOrderCancelRequest struct {
    model.Params

    // 物流宝订单编号
    wlbOrderCode   string 

}

func NewTaobaoWlbOrderCancelRequest() *TaobaoWlbOrderCancelRequest{
    return &TaobaoWlbOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderCancelRequest) GetApiMethodName() string {
    return "taobao.wlb.order.cancel"
}

func (r TaobaoWlbOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderCancelRequest) SetWlbOrderCode(wlbOrderCode string) error {
    r.wlbOrderCode = wlbOrderCode
    r.Set("wlb_order_code", wlbOrderCode)
    return nil
}

func (r TaobaoWlbOrderCancelRequest) GetWlbOrderCode() string {
    return r.wlbOrderCode
}

