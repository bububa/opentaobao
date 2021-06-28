package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单仓库路由信息 APIResponse
taobao.wms.order.warehouse.route.get

获取订单仓库路由信息
*/
type TaobaoWmsOrderWarehouseRouteGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wms_order_warehouse_route_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品列表
    
    Items   []OrderWarehouseRouteGetItems `json:"items,omitempty" xml:"