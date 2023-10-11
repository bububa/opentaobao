package wdk

// NetworkRouteDto 结构体
type NetworkRouteDto struct {
	// 节点信息
	Nodes []LogisticsNodeFullInfo `json:"nodes,omitempty" xml:"nodes>logistics_node_full_info,omitempty"`
	// 线路ID
	RouteId string `json:"route_id,omitempty" xml:"route_id,omitempty"`
}
