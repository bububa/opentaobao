package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取订单仓库路由信息 APIResponse
taobao.wms.order.warehouse.route.get

获取订单仓库路由信息
*/
type TaobaoWmsOrderWarehouseRouteGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWmsOrderWarehouseRouteGetResponse `json:"wms_order_warehouse_route_get_response,omitempty"` 
    TaobaoWmsOrderWarehouseRouteGetResponse
}

/* model for simplify = false
type TaobaoWmsOrderWarehouseRouteGetResponse struct {

    // 商品列表
    
    Items  struct {
        OrderWarehouseRouteGetItems  []OrderWarehouseRouteGetItems `json:"order_warehouse_route_get_items,omitempty"`
    } `json:"items,omitempty"`
    

    // 是否成功
    
    WlSuccess   bool `json:"wl_success,omitempty"`
    

    // 错误信息
    
    WlErrorCode   string `json:"wl_error_code,omitempty"`
    

    // 错误信息
    
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`
    

    // 订单编号
    
    OrderCode   string `json:"order_code,omitempty"`
    

}
*/

type TaobaoWmsOrderWarehouseRouteGetResponse struct {

    // 商品列表
    Items   []OrderWarehouseRouteGetItems `json:"items,omitempty"`

    // 是否成功
    WlSuccess   bool `json:"wl_success,omitempty"`

    // 错误信息
    WlErrorCode   string `json:"wl_error_code,omitempty"`

    // 错误信息
    WlErrorMsg   string `json:"wl_error_msg,omitempty"`

    // 订单编号
    OrderCode   string `json:"order_code,omitempty"`

}
