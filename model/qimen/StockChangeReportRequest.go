package qimen

// StockChangeReportRequest 结构体
type StockChangeReportRequest struct {
	// item
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockchangeReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
