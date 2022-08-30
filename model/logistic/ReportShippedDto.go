package logistic

// ReportShippedDto 结构体
type ReportShippedDto struct {
	// shipment order id
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
	// type of report dispatch
	ShippedType string `json:"shipped_type,omitempty" xml:"shipped_type,omitempty"`
}
