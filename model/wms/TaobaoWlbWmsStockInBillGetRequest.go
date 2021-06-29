package wms

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取入库单信息 API请求
taobao.wlb.wms.stock.in.bill.get

获取入库单信息
*/
type TaobaoWlbWmsStockInBillGetRequest struct {
    model.Params
    // ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    orderCode   string
    // 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
    cnOrderCode   string
}

// 初始化TaobaoWlbWmsStockInBillGetRequest对象
func NewTaobaoWlbWmsStockInBillGetRequest() *TaobaoWlbWmsStockInBillGetRequest{
    return &TaobaoWlbWmsStockInBillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockInBillGetRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.stock.in.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockInBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockInBillGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWlbWmsStockInBillGetRequest) GetOrderCode() string {
    return r.orderCode
}
// CnOrderCode Setter
// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockInBillGetRequest) SetCnOrderCode(cnOrderCode string) error {
    r.cnOrderCode = cnOrderCode
    r.Set("cn_order_code", cnOrderCode)
    return nil
}

// CnOrderCode Getter
func (r TaobaoWlbWmsStockInBillGetRequest) GetCnOrderCode() string {
    return r.cnOrderCode
}
