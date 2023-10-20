package btrip

import (
	"sync"
)

// RouteInfoRs 结构体
type RouteInfoRs struct {
	// 组成当前线路的形成列表
	JourneyList []JourneyRs `json:"journey_list,omitempty" xml:"journey_list>journey_rs,omitempty"`
}

var poolRouteInfoRs = sync.Pool{
	New: func() any {
		return new(RouteInfoRs)
	},
}

// GetRouteInfoRs() 从对象池中获取RouteInfoRs
func GetRouteInfoRs() *RouteInfoRs {
	return poolRouteInfoRs.Get().(*RouteInfoRs)
}

// ReleaseRouteInfoRs 释放RouteInfoRs
func ReleaseRouteInfoRs(v *RouteInfoRs) {
	v.JourneyList = v.JourneyList[:0]
	poolRouteInfoRs.Put(v)
}
