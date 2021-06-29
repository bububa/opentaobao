package feedflow

// LaunchPeriodDTO 
type LaunchPeriodDTO struct {
    // 时间
    TimeSpanList   []TimeSpanDTO `json:"time_span_list,omitempty" xml:"time_span_list>time_span_dto,omitempty"`
}
