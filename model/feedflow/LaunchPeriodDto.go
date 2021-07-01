package feedflow

// LaunchPeriodDto 结构体
type LaunchPeriodDto struct {
	// 时间
	TimeSpanList []TimeSpanDto `json:"time_span_list,omitempty" xml:"time_span_list>time_span_dto,omitempty"`
}
