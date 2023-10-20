package btrip

import (
	"sync"
)

// RouteVo 结构体
type RouteVo struct {
	// 来自城市
	FromCity string `json:"from_city,omitempty" xml:"from_city,omitempty"`
	// 到达城市
	ToCity string `json:"to_city,omitempty" xml:"to_city,omitempty"`
}

var poolRouteVo = sync.Pool{
	New: func() any {
		return new(RouteVo)
	},
}

// GetRouteVo() 从对象池中获取RouteVo
func GetRouteVo() *RouteVo {
	return poolRouteVo.Get().(*RouteVo)
}

// ReleaseRouteVo 释放RouteVo
func ReleaseRouteVo(v *RouteVo) {
	v.FromCity = ""
	v.ToCity = ""
	poolRouteVo.Put(v)
}
