package openim

// RoamingMessageItem 结构体
type RoamingMessageItem struct {
	// 节点类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
