package btrip

// RouteInfoRS 
type RouteInfoRS struct {
    // 组成当前线路的形成列表
    JourneyList   []JourneyRS `json:"journey_list,omitempty" xml:"journey_list>journey_rs,omitempty"`
}
