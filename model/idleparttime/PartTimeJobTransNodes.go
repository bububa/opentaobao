package idleparttime

// PartTimeJobTransNodes 结构体
type PartTimeJobTransNodes struct {
	// 节点描述
	NodeDescription string `json:"node_description,omitempty" xml:"node_description,omitempty"`
	// 节点创建时间
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}
