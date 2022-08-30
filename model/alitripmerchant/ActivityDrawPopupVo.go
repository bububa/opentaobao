package alitripmerchant

// ActivityDrawPopupVo 结构体
type ActivityDrawPopupVo struct {
	// 弹窗点击文本
	PopupDoText string `json:"popup_do_text,omitempty" xml:"popup_do_text,omitempty"`
	// 标题
	PopupTitle string `json:"popup_title,omitempty" xml:"popup_title,omitempty"`
	// 文本
	PopupSubText string `json:"popup_sub_text,omitempty" xml:"popup_sub_text,omitempty"`
	// 头部标题
	TopTitle string `json:"top_title,omitempty" xml:"top_title,omitempty"`
	// 活动id
	OfferId string `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 弹窗文本
	PopupExitText string `json:"popup_exit_text,omitempty" xml:"popup_exit_text,omitempty"`
	// 文本
	PopupText string `json:"popup_text,omitempty" xml:"popup_text,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 抽奖类型
	PopupType int64 `json:"popup_type,omitempty" xml:"popup_type,omitempty"`
}
