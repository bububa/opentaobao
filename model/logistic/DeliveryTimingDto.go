package logistic

// DeliveryTimingDto 结构体
type DeliveryTimingDto struct {
	// 预估时效，这部分ISV直接展示，不要做改动，ISP会变文案和时效展示
	DeliveryPeriod string `json:"delivery_period,omitempty" xml:"delivery_period,omitempty"`
}
