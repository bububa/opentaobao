package wdk

// RouteNodeDto 结构体
type RouteNodeDto struct {
	// 节点序号
	NodeIndex int64 `json:"node_index,omitempty" xml:"node_index,omitempty"`
	// 节点类型
	NodeType string `json:"node_type,omitempty" xml:"node_type,omitempty"`
	// 节点名称
	NodeCode string `json:"node_code,omitempty" xml:"node_code,omitempty"`
}
