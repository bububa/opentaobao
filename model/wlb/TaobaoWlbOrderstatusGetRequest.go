package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流宝订单流转状态查询 APIRequest
taobao.wlb.orderstatus.get

根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
*/
type TaobaoWlbOrderstatusGetRequest struct {
    model.Params

    // 物流宝订单编码
    orderCode   string 

}

func NewTaobaoWlbOrderstatusGetRequest() *TaobaoWlbOrderstatusGetRequest{
    return &TaobaoWlbOrderstatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbOrderstatusGetRequest) GetApiMethodName() string {
    return "taobao.wlb.orderstatus.get"
}

func (r TaobaoWlbOrderstatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbOrderstatusGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbOrderstatusGetRequest) GetOrderCode() string {
    return r.orderCode
}

