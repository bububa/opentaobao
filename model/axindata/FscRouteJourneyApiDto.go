package axindata

// FscRouteJourneyApiDto 结构体
type FscRouteJourneyApiDto struct {
	// 日程列表
	RouteScheduleList []FscRouteScheduleApiDto `json:"route_schedule_list,omitempty" xml:"route_schedule_list>fsc_route_schedule_api_dto,omitempty"`
	// 行程编号
	JourneyCode string `json:"journey_code,omitempty" xml:"journey_code,omitempty"`
	// 行程名称
	JourneyName string `json:"journey_name,omitempty" xml:"journey_name,omitempty"`
	// 行程描述
	JourneyDesc string `json:"journey_desc,omitempty" xml:"journey_desc,omitempty"`
	// 出发城市名称
	StartCityName string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty"`
	// 出发城市Id
	StartCityId int64 `json:"start_city_id,omitempty" xml:"start_city_id,omitempty"`
}
