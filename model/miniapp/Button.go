package miniapp

// Button 结构体
type Button struct {
	// 文案
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 跳转链接
	TargetUrl string `json:"target_url,omitempty" xml:"target_url,omitempty"`
}
