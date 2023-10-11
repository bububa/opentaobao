package simba

// LaunchPeriodVo 结构体
type LaunchPeriodVo struct {
	// 第x天各时段的折扣情况
	TimeSpanList []TimeSpanQueryResVo `json:"time_span_list,omitempty" xml:"time_span_list>time_span_query_res_vo,omitempty"`
	// 本周的第x天
	DayOfWeek int64 `json:"day_of_week,omitempty" xml:"day_of_week,omitempty"`
}
