package wdk

import (
	"sync"
)

// NetworkRouteDto 结构体
type NetworkRouteDto struct {
	// 节点信息
	Nodes []LogisticsNodeFullInfo `json:"nodes,omitempty" xml:"nodes>logistics_node_full_info,omitempty"`
	// 线路ID
	RouteId string `json:"route_id,omitempty" xml:"route_id,omitempty"`
}

var poolNetworkRouteDto = sync.Pool{
	New: func() any {
		return new(NetworkRouteDto)
	},
}

// GetNetworkRouteDto() 从对象池中获取NetworkRouteDto
func GetNetworkRouteDto() *NetworkRouteDto {
	return poolNetworkRouteDto.Get().(*NetworkRouteDto)
}

// ReleaseNetworkRouteDto 释放NetworkRouteDto
func ReleaseNetworkRouteDto(v *NetworkRouteDto) {
	v.Nodes = v.Nodes[:0]
	v.RouteId = ""
	poolNetworkRouteDto.Put(v)
}
