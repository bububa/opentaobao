package perfect

// LoadPackageOrderDetailRequest 结构体
type LoadPackageOrderDetailRequest struct {
	// 销量数量
	SalesQuantity string `json:"sales_quantity,omitempty" xml:"sales_quantity,omitempty"`
	// 货品
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 销量单位
	SalesUnit string `json:"sales_unit,omitempty" xml:"sales_unit,omitempty"`
	// 库存数量
	StockQuantity string `json:"stock_quantity,omitempty" xml:"stock_quantity,omitempty"`
	// 扩展
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 出库子单
	OutboundOrderDetailCode string `json:"outbound_order_detail_code,omitempty" xml:"outbound_order_detail_code,omitempty"`
	// 库存单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
}
