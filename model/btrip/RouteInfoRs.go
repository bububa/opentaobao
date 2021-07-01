package btrip

// RouteInfoRs 结构体
type RouteInfoRs struct {
	// 组成当前线路的形成列表
	JourneyList []JourneyRs `json:"journey_list,omitempty" xml:"journey_list>journey_rs,omitempty"`
}
