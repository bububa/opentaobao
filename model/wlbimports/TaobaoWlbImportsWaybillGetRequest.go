package wlbimports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取运单信息 APIRequest
taobao.wlb.imports.waybill.get

一般进口商家，获取订单的电子面单链接地址
*/
type TaobaoWlbImportsWaybillGetRequest struct {
    model.Params

    // 物流订单号
    orderCode   string 

}

func NewTaobaoWlbImportsWaybillGetRequest() *TaobaoWlbImportsWaybillGetRequest{
    return &TaobaoWlbImportsWaybillGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbImportsWaybillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.imports.waybill.get"
}

func (r TaobaoWlbImportsWaybillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbImportsWaybillGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbImportsWaybillGetRequest) GetOrderCode() string {
    return r.orderCode
}

