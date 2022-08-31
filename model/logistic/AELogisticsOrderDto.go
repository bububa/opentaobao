package logistic

// AELogisticsOrderDto 结构体
type AELogisticsOrderDto struct {
	// Shipment order id created. AE will save relationship with logistics provider&#39;s shipment order_id
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
}
