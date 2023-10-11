package product

// GroupItem 结构体
type GroupItem struct {
	// 分组名称
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 分组id
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}
