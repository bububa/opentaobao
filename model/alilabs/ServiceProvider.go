package alilabs

// ServiceProvider 结构体
type ServiceProvider struct {
	// 图片地址
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 提供商名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
