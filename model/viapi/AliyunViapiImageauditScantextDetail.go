package viapi

// AliyunViapiImageauditScantextDetail 结构体
type AliyunViapiImageauditScantextDetail struct {
	// 命中该风险的上下文信息
	Contexts []Context `json:"contexts,omitempty" xml:"contexts>context,omitempty"`
	// 文本命中风险的分类
	Label string `json:"label,omitempty" xml:"label,omitempty"`
}
