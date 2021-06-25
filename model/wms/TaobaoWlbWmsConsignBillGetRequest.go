package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取销售订单发货信息 APIRequest
taobao.wlb.wms.consign.bill.get

获取销售订单发货信息
*/
type TaobaoWlbWmsConsignBillGetRequest struct {
    model.Params

    // 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
    cnOrderCode   string 

    // ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
    orderCode   string 

}

func NewTaobaoWlbWmsConsignBillGetRequest() *TaobaoWlbWmsConsignBillGetRequest{
    return &TaobaoWlbWmsConsignBillGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsConsignBillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.consign.bill.get"
}

func (r TaobaoWlbWmsConsignBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsConsignBillGetRequest) SetCnOrderCode(cnOrderCode string) error {
    r.cnOrderCode = cnOrderCode
    r.Set("cn_order_code", cnOrderCode)
    return nil
}

func (r TaobaoWlbWmsConsignBillGetRequest) GetCnOrderCode() string {
    return r.cnOrderCode
}

func (r *TaobaoWlbWmsConsignBillGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbWmsConsignBillGetRequest) GetOrderCode() string {
    return r.orderCode
}

