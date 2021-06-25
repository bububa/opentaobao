package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝订单已发货通知接口 APIRequest
taobao.wlb.order.consign

如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
*/
type TaobaoWlbOrderConsignRequest struct {
    model.Params

    // 物流宝订单编号
    wlbOrderCode   string 

}

func NewTaobaoWlbOrderConsignRequest() *TaobaoWlbOrderConsignRequest{
    return &TaobaoWlbOrderConsignRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderConsignRequest) GetApiMethodName() string {
    return "taobao.wlb.order.consign"
}

func (r TaobaoWlbOrderConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderConsignRequest) SetWlbOrderCode(wlbOrderCode string) error {
    r.wlbOrderCode = wlbOrderCode
    r.Set("wlb_order_code", wlbOrderCode)
    return nil
}

func (r TaobaoWlbOrderConsignRequest) GetWlbOrderCode() string {
    return r.wlbOrderCode
}

