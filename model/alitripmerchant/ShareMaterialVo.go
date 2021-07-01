package alitripmerchant

// ShareMaterialVo 结构体
type ShareMaterialVo struct {
	// 分享的文案
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 分享的图片
	Picture string `json:"picture,omitempty" xml:"picture,omitempty"`
}
