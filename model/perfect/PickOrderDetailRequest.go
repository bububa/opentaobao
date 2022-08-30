package perfect

// PickOrderDetailRequest 结构体
type PickOrderDetailRequest struct {
	// 1
	Barcodes []string `json:"barcodes,omitempty" xml:"barcodes>string,omitempty"`
	// 真实拣货数量
	RealPickQuantity string `json:"real_pick_quantity,omitempty" xml:"real_pick_quantity,omitempty"`
	// 拣货完成时间
	PickFinishTime string `json:"pick_finish_time,omitempty" xml:"pick_finish_time,omitempty"`
	// 1
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 1
	PickUnit string `json:"pick_unit,omitempty" xml:"pick_unit,omitempty"`
	// 1
	PlanStockQuantity string `json:"plan_stock_quantity,omitempty" xml:"plan_stock_quantity,omitempty"`
	// 1
	RealStockQuantity string `json:"real_stock_quantity,omitempty" xml:"real_stock_quantity,omitempty"`
	// 1
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 1
	PickOrderDetailCode string `json:"pick_order_detail_code,omitempty" xml:"pick_order_detail_code,omitempty"`
	// 1
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 1
	PlanPickQuantity string `json:"plan_pick_quantity,omitempty" xml:"plan_pick_quantity,omitempty"`
	// 1
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 1
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
}
