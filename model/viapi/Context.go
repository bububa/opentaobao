package viapi

// Context 结构体
type Context struct {
	// 检测文本命中的风险内容上下文内容
	Context string `json:"context,omitempty" xml:"context,omitempty"`
}
