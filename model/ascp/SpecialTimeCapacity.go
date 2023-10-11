package ascp

// SpecialTimeCapacity 结构体
type SpecialTimeCapacity struct {
	// 开始时间
	StartHour string `json:"start_hour,omitempty" xml:"start_hour,omitempty"`
	// 结束时间
	ToHour string `json:"to_hour,omitempty" xml:"to_hour,omitempty"`
	// 产能
	Capacity string `json:"capacity,omitempty" xml:"capacity,omitempty"`
}
