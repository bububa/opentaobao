package alihouse

// ProjectOrderDto 结构体
type ProjectOrderDto struct {
	// 交易商品列表
	TradeItemList []ProjectTradeItemOrderDto `json:"trade_item_list,omitempty" xml:"trade_item_list>project_trade_item_order_dto,omitempty"`
	// 模块排序信息
	ModuleOrder []ModuleTypeOrderDto `json:"module_order,omitempty" xml:"module_order>module_type_order_dto,omitempty"`
	// 楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 排序号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
}
