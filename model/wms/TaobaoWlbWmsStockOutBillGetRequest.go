package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单号获取单个出库单发货信息 API请求
taobao.wlb.wms.stock.out.bill.get

通过订单号获取单个出库单发货信息
*/
type TaobaoWlbWmsStockOutBillGetAPIRequest struct {
    model.Params
    // ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    _orderCode   string
    // 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    _cnOrderCode   string
}

// 初始化TaobaoWlbWmsStockOutBillGetAPIRequest对象
func NewTaobaoWlbWmsStockOutBillGetRequest() *TaobaoWlbWmsStockOutBillGetAPIRequest{
    return &TaobaoWlbWmsStockOutBillGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockOutBillGetAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.out.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockOutBillGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockOutBillGetAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsStockOutBillGetAPIRequest) GetOrderCode() string {
    return r._orderCode
}
// CnOrderCode Setter
// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockOutBillGetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
    r._cnOrderCode = _cnOrderCode
    r.Set("cn_order_code", _cnOrderCode)
    return nil
}

// CnOrderCode Getter
func (r TaobaoWlbWmsStockOutBillGetAPIRequest) GetCnOrderCode() string {
    return r._cnOrderCode
}
