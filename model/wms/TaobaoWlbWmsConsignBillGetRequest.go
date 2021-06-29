package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取销售订单发货信息 API请求
taobao.wlb.wms.consign.bill.get

获取销售订单发货信息
*/
type TaobaoWlbWmsConsignBillGetRequest struct {
    model.Params
    // 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
    _cnOrderCode   string
    // ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
    _orderCode   string
}

// 初始化TaobaoWlbWmsConsignBillGetRequest对象
func NewTaobaoWlbWmsConsignBillGetRequest() *TaobaoWlbWmsConsignBillGetRequest{
    return &TaobaoWlbWmsConsignBillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsConsignBillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.consign.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsConsignBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CnOrderCode Setter
// 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
func (r *TaobaoWlbWmsConsignBillGetRequest) SetCnOrderCode(_cnOrderCode string) error {
    r._cnOrderCode = _cnOrderCode
    r.Set("cn_order_code", _cnOrderCode)
    return nil
}

// CnOrderCode Getter
func (r TaobaoWlbWmsConsignBillGetRequest) GetCnOrderCode() string {
    return r._cnOrderCode
}
// OrderCode Setter
// ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
func (r *TaobaoWlbWmsConsignBillGetRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsConsignBillGetRequest) GetOrderCode() string {
    return r._orderCode
}
