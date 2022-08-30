package aedropshiper

// AeOrderLogisticsInfo 结构体
type AeOrderLogisticsInfo struct {
	// Logistics tracking number
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// Logistics Services
	LogisticsService string `json:"logistics_service,omitempty" xml:"logistics_service,omitempty"`
}
