package axindata

// FscRouteScheduleApiDto 结构体
type FscRouteScheduleApiDto struct {
	// 日程明细列表
	ScheduleDetailList []FscRouteScheduleDetailApiDto `json:"schedule_detail_list,omitempty" xml:"schedule_detail_list>fsc_route_schedule_detail_api_dto,omitempty"`
	// 日程名称
	ScheduleName string `json:"schedule_name,omitempty" xml:"schedule_name,omitempty"`
	// 日程描述
	ScheduleDesc string `json:"schedule_desc,omitempty" xml:"schedule_desc,omitempty"`
	// 日程天
	ScheduleDay int64 `json:"schedule_day,omitempty" xml:"schedule_day,omitempty"`
}
