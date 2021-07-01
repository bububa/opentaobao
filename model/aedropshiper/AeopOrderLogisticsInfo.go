package aedropshiper

// AeopOrderLogisticsInfo 结构体
type AeopOrderLogisticsInfo struct {
	// 物流追踪号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 物流服务
	LogisticsService string `json:"logistics_service,omitempty" xml:"logistics_service,omitempty"`
}
