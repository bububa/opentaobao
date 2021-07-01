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
type TaobaoWlbWmsConsignBillGetAPIRequest struct {
    model.Params
    // 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
    _cnOrderCode   string
    // ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
    _orderCode   string
}

// 初始化TaobaoWlbWmsConsignBillGetAPIRequest对象
func NewTaobaoWlbWmsConsignBillGetRequest() *TaobaoWlbWmsConsignBillGetAPIRequest{
    return &TaobaoWlbWmsConsignBillGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsConsignBillGetAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.consign.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsConsignBillGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CnOrderCode Setter
// 菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空
func (r *TaobaoWlbWmsConsignBillGetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
    r._cnOrderCode = _cnOrderCode
    r.Set("cn_order_code", _cnOrderCode)
    return nil
}

// CnOrderCode Getter
func (r TaobaoWlbWmsConsignBillGetAPIRequest) GetCnOrderCode() string {
    return r._cnOrderCode
}
// OrderCode Setter
// ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空
func (r *TaobaoWlbWmsConsignBillGetAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsConsignBillGetAPIRequest) GetOrderCode() string {
    return r._orderCode
}
