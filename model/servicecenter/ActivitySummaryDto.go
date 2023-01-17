package servicecenter

// ActivitySummaryDto 结构体
type ActivitySummaryDto struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间（设置）
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 中止时间
	TerminateTime string `json:"terminate_time,omitempty" xml:"terminate_time,omitempty"`
}
