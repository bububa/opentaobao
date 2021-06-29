package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单仓库路由信息 API请求
taobao.wms.order.warehouse.route.get

获取订单仓库路由信息
*/
type TaobaoWmsOrderWarehouseRouteGetRequest struct {
    model.Params
    // 订单编号
    orderCode   string
}

// 初始化TaobaoWmsOrderWarehouseRouteGetRequest对象
func NewTaobaoWmsOrderWarehouseRouteGetRequest() *TaobaoWmsOrderWarehouseRouteGetRequest{
    return &TaobaoWmsOrderWarehouseRouteGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWmsOrderWarehouseRouteGetRequest) GetApiMethodName() string {
    return "taobao.wms.order.warehouse.route.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWmsOrderWarehouseRouteGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 订单编号
func (r *TaobaoWmsOrderWarehouseRouteGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWmsOrderWarehouseRouteGetRequest) GetOrderCode() string {
    return r.orderCode
}
