package charity

// SumUserCharityTimeTopApiHsfResponse 结构体
type SumUserCharityTimeTopApiHsfResponse struct {
	// 公益时：分钟
	CharityTimeMinutes int64 `json:"charity_time_minutes,omitempty" xml:"charity_time_minutes,omitempty"`
	// 公益时：小时
	CharityTimeHours float64 `json:"charity_time_hours,omitempty" xml:"charity_time_hours,omitempty"`
}
