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
type TaobaoWlbWmsStockOutBillGetRequest struct {
    model.Params
    // ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    orderCode   string
    // 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    cnOrderCode   string
}

// 初始化TaobaoWlbWmsStockOutBillGetRequest对象
func NewTaobaoWlbWmsStockOutBillGetRequest() *TaobaoWlbWmsStockOutBillGetRequest{
    return &TaobaoWlbWmsStockOutBillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockOutBillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.out.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockOutBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockOutBillGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsStockOutBillGetRequest) GetOrderCode() string {
    return r.orderCode
}
// CnOrderCode Setter
// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockOutBillGetRequest) SetCnOrderCode(cnOrderCode string) error {
    r.cnOrderCode = cnOrderCode
    r.Set("cn_order_code", cnOrderCode)
    return nil
}

// CnOrderCode Getter
func (r TaobaoWlbWmsStockOutBillGetRequest) GetCnOrderCode() string {
    return r.cnOrderCode
}
