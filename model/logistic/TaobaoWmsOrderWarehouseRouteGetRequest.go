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
type TaobaoWmsOrderWarehouseRouteGetAPIRequest struct {
    model.Params
    // 订单编号
    _orderCode   string
}

// 初始化TaobaoWmsOrderWarehouseRouteGetAPIRequest对象
func NewTaobaoWmsOrderWarehouseRouteGetRequest() *TaobaoWmsOrderWarehouseRouteGetAPIRequest{
    return &TaobaoWmsOrderWarehouseRouteGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWmsOrderWarehouseRouteGetAPIRequest) GetApiMethodName() string {
    return "taobao.wms.order.warehouse.route.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWmsOrderWarehouseRouteGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 订单编号
func (r *TaobaoWmsOrderWarehouseRouteGetAPIRequest) SetOrderCode(_orderCode string) error {
    r._orderCode = _orderCode
    r.Set("order_code", _orderCode)
    return nil
}

// OrderCode Getter
func (r TaobaoWmsOrderWarehouseRouteGetAPIRequest) GetOrderCode() string {
    return r._orderCode
}
