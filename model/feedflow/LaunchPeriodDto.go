package feedflow

// LaunchPeriodDto 
type LaunchPeriodDto struct {
    // 列表
    TimeSpanList   []TimeSpanDto `json:"time_span_list,omitempty" xml:"time_span_list>time_span_dto,omitempty"`
}
