package tbk

// TaobaotbkactivityinfogetData 结构体
type TaobaotbkactivityinfogetData struct {
	// 【本地化】微信推广二维码地址
	Wxqrcodeurl string `json:"wx_qrcode_url,omitempty" xml:"wx_qrcode_url,omitempty"`
	// 淘客推广长链
	Clickurl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 淘客推广短链
	Shortclickurl string `json:"short_click_url,omitempty" xml:"short_click_url,omitempty"`
	// 投放平台, 1-PC 2-无线
	Terminaltype string `json:"terminal_type,omitempty" xml:"terminal_type,omitempty"`
	// 物料素材下载地址
	Materialossurl string `json:"material_oss_url,omitempty" xml:"material_oss_url,omitempty"`
	// 会场名称
	Pagename string `json:"page_name,omitempty" xml:"page_name,omitempty"`
	// 活动开始时间
	Pagestarttime string `json:"page_start_time,omitempty" xml:"page_start_time,omitempty"`
	// 活动结束时间
	Pageendtime string `json:"page_end_time,omitempty" xml:"page_end_time,omitempty"`
	// 【本地化】微信小程序路径
	Wxminiprogrampath string `json:"wx_miniprogram_path,omitempty" xml:"wx_miniprogram_path,omitempty"`
}
