package alihealthoutflow

// QuickAppCardInfoVo 结构体
type QuickAppCardInfoVo struct {
	// 背景图
	BgImg string `json:"bg_img,omitempty" xml:"bg_img,omitempty"`
	// 背景图右侧图标
	BgRightIcon string `json:"bg_right_icon,omitempty" xml:"bg_right_icon,omitempty"`
	// logo图标
	LogoImg string `json:"logo_img,omitempty" xml:"logo_img,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 标题颜色
	TitleColor string `json:"title_color,omitempty" xml:"title_color,omitempty"`
	// 标题字体大小
	TitleSize string `json:"title_size,omitempty" xml:"title_size,omitempty"`
	// 副标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 副标题颜色
	SubTitleColor string `json:"sub_title_color,omitempty" xml:"sub_title_color,omitempty"`
	// 副标题大小
	SubTitleSize string `json:"sub_title_size,omitempty" xml:"sub_title_size,omitempty"`
	// 按钮背景颜色
	ButtonBgColor string `json:"button_bg_color,omitempty" xml:"button_bg_color,omitempty"`
	// 按钮文案
	ButtonText string `json:"button_text,omitempty" xml:"button_text,omitempty"`
	// 按钮字体颜色
	ButtonTextColor string `json:"button_text_color,omitempty" xml:"button_text_color,omitempty"`
	// 按钮字体大小
	ButtonTextSize string `json:"button_text_size,omitempty" xml:"button_text_size,omitempty"`
	// 跳转链接
	JumpUrl string `json:"jump_url,omitempty" xml:"jump_url,omitempty"`
}
