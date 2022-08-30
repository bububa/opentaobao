package perfect

// PickOutboundOrderDetailRequest 结构体
type PickOutboundOrderDetailRequest struct {
	// 1
	Barcodes []string `json:"barcodes,omitempty" xml:"barcodes>string,omitempty"`
	// 1
	ItemPicUrl string `json:"item_pic_url,omitempty" xml:"item_pic_url,omitempty"`
	// 1
	ItemType string `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 1
	PlanSalesQuantity string `json:"plan_sales_quantity,omitempty" xml:"plan_sales_quantity,omitempty"`
	// 1
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 1
	PlanStockQuantity string `json:"plan_stock_quantity,omitempty" xml:"plan_stock_quantity,omitempty"`
	// 1
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 1
	InterceptStrategy string `json:"intercept_strategy,omitempty" xml:"intercept_strategy,omitempty"`
	// 1
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 1
	SalesUnit string `json:"sales_unit,omitempty" xml:"sales_unit,omitempty"`
	// 1
	Cancelled string `json:"cancelled,omitempty" xml:"cancelled,omitempty"`
	// 1
	ChannelName string `json:"channel_name,omitempty" xml:"channel_name,omitempty"`
	// 1
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 1
	OutboundOrderDetailCode string `json:"outbound_order_detail_code,omitempty" xml:"outbound_order_detail_code,omitempty"`
}
