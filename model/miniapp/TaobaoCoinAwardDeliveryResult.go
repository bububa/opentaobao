package miniapp

// TaobaoCoinAwardDeliveryResult 结构体
type TaobaoCoinAwardDeliveryResult struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 副标题
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 按钮素材
	Button *Button `json:"button,omitempty" xml:"button,omitempty"`
	// 图片素材
	Image *Image `json:"image,omitempty" xml:"image,omitempty"`
}
