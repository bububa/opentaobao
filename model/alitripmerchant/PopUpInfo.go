package alitripmerchant

// PopUpInfo 结构体
type PopUpInfo struct {
	// 弹屏图片
	PopUpUrl string `json:"pop_up_url,omitempty" xml:"pop_up_url,omitempty"`
	// 弹屏跳转链接
	PopUpRedirectUrl string `json:"pop_up_redirect_url,omitempty" xml:"pop_up_redirect_url,omitempty"`
}
