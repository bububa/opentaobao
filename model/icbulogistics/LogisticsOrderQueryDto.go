package icbulogistics

// LogisticsOrderQueryDto 结构体
type LogisticsOrderQueryDto struct {
	// 物流单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
}
