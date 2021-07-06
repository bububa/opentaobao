package logistic

// OrderWarehouseRouteGetItem 结构体
type OrderWarehouseRouteGetItem struct {
	// ERP订单明细编码或者子交易单号
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 菜鸟订单编码 当订单路由到菜鸟仓发货时，生成此单号。等待路由仓或由商家仓发货的订单，此单号为空。格式LBX+数字
	CnOrderCode string `json:"cn_order_code,omitempty" xml:"cn_order_code,omitempty"`
	// 订单路由状态 状态值： WAIT_ROUTE 待路由仓 ROUTE_TO_CN 路由到菜鸟仓发货 ROUTE_TO_ERP 路由到商家仓发货。 STOP_ROUTE 终止分仓，如订单已取消时，不再发货。 注：待路由仓状态表示未做路由，不确定由哪个仓库发货，可与5分钟后再次查询
	RoutStatus string `json:"rout_status,omitempty" xml:"rout_status,omitempty"`
	// 仓库编码 当订单路由到菜鸟仓发货时输出菜鸟仓编码。等待路由仓或由商家仓发货的订单，此内容为空。
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 通知仓库此订单明细的商品应发数量
	ItemQty int64 `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
}
