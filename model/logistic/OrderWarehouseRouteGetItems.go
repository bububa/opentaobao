package logistic

// OrderWarehouseRouteGetItems 
/* model for simplify = false
type OrderWarehouseRouteGetItems struct {

    // 商品信息
    
    Item  *struct {
        OrderWarehouseRouteGetItem  *OrderWarehouseRouteGetItem `json:"order_warehouse_route_get_item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

// OrderWarehouseRouteGetItems 
type OrderWarehouseRouteGetItems struct {

    // 商品信息
    Item   *OrderWarehouseRouteGetItem `json:"item,omitempty"`

}
