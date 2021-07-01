package openim

// MessageItem 结构体
type MessageItem struct {
	// 节点类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 节点值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
