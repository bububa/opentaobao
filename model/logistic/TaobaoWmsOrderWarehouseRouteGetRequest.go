package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单仓库路由信息 APIRequest
taobao.wms.order.warehouse.route.get

获取订单仓库路由信息
*/
type TaobaoWmsOrderWarehouseRouteGetRequest struct {
    model.Params

    // 订单编号
    orderCode   string 

}

func NewTaobaoWmsOrderWarehouseRouteGetRequest() *TaobaoWmsOrderWarehouseRouteGetRequest{
    return &TaobaoWmsOrderWarehouseRouteGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWmsOrderWarehouseRouteGetRequest) GetApiMethodName() string {
    return "taobao.wms.order.warehouse.route.get"
}

func (r TaobaoWmsOrderWarehouseRouteGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWmsOrderWarehouseRouteGetRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("order_code", orderCode)
    return nil
}

func (r TaobaoWmsOrderWarehouseRouteGetRequest) GetOrderCode() string {
    return r.orderCode
}

