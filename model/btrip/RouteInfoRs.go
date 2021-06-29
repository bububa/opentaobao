package btrip

// RouteInfoRs 
type RouteInfoRs struct {

    // 组成当前线路的形成列表
    
    JourneyList   []JourneyRs `json:"journey_list,omitempty" xml:"journey_list,omitempty"`
    

}
