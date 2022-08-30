package tmallgenie

// LivingGenericEventDto 结构体
type LivingGenericEventDto struct {
	// 事件内容
	Payload string `json:"payload,omitempty" xml:"payload,omitempty"`
	// 事件命名空间
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// 命名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
