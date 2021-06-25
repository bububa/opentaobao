package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据物流宝订单编号查询物流宝订单概要信息 APIRequest
taobao.wlb.wlborder.get

根据物流宝订单编号查询物流宝订单概要信息
*/
type TaobaoWlbWlborderGetRequest struct {
    model.Params

    // 物流宝订单编码
    wlbOrderCode   string 

}

func NewTaobaoWlbWlborderGetRequest() *TaobaoWlbWlborderGetRequest{
    return &TaobaoWlbWlborderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWlborderGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wlborder.get"
}

func (r TaobaoWlbWlborderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWlborderGetRequest) SetWlbOrderCode(wlbOrderCode string) error {
    r.wlbOrderCode = wlbOrderCode
    r.Set("wlb_order_code", wlbOrderCode)
    return nil
}

func (r TaobaoWlbWlborderGetRequest) GetWlbOrderCode() string {
    return r.wlbOrderCode
}

