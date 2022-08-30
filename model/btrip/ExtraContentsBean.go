package btrip

// ExtraContentsBean 结构体
type ExtraContentsBean struct {
	// 说明内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 内容标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}
