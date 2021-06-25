package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取销退收货信息 APIRequest
taobao.wlb.wms.return.bill.get

通过订单号获取单个销退单收货信息。
*/
type TaobaoWlbWmsReturnBillGetRequest struct {
    model.Params

    // ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    orderCode   string 

    // 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    cnOrderCode   string 

}

func NewTaobaoWlbWmsReturnBillGetRequest() *TaobaoWlbWmsReturnBillGetRequest{
    return &TaobaoWlbWmsReturnBillGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsReturnBillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.return.bill.get"
}

func (r TaobaoWlbWmsReturnBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsReturnBillGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWlbWmsReturnBillGetRequest) GetOrderCode() string {
    return r.orderCode
}

func (r *TaobaoWlbWmsReturnBillGetRequest) SetCnOrderCode(cnOrderCode string) error {
    r.cnOrderCode = cnOrderCode
    r.Set("cn_order_code", cnOrderCode)
    return nil
}

func (r TaobaoWlbWmsReturnBillGetRequest) GetCnOrderCode() string {
    return r.cnOrderCode
}

