package wdk

// WorkOrderUnits 结构体
type WorkOrderUnits struct {
	// 作业子单列表
	WorkOrderUnitContents []WorkOrderUnitContents `json:"work_order_unit_contents,omitempty" xml:"work_order_unit_contents>work_order_unit_contents,omitempty"`
	// 门店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 履约单号/订单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
	// 作业单最早送达时间
	EarliestArrivalTime string `json:"earliest_arrival_time,omitempty" xml:"earliest_arrival_time,omitempty"`
	// 作业单最晚送达时间
	LatestArriveTime string `json:"latest_arrive_time,omitempty" xml:"latest_arrive_time,omitempty"`
	// 订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 额外属性
	ExtMap *ExtMap `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 收货人对象
	Customer *Customer `json:"customer,omitempty" xml:"customer,omitempty"`
	// 订单来源
	SourceFrom int64 `json:"source_from,omitempty" xml:"source_from,omitempty"`
}
