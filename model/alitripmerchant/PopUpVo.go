package alitripmerchant

// PopUpVo 结构体
type PopUpVo struct {
	// 弹屏跳转链接
	PopUpRedirectUrl string `json:"pop_up_redirect_url,omitempty" xml:"pop_up_redirect_url,omitempty"`
	// 弹屏图片
	PopUpUrl string `json:"pop_up_url,omitempty" xml:"pop_up_url,omitempty"`
	// 活动id
	OfferId int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
}
