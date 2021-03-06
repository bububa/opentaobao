package servicecenter

// CsSchedulingWrapper 结构体
type CsSchedulingWrapper struct {
	// 按天排班信息
	CsSchedulings []CsScheduling `json:"cs_schedulings,omitempty" xml:"cs_schedulings>cs_scheduling,omitempty"`
	// 排班记录条数（按天计算）
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
