package tbk

import (
	"sync"
)

// TaobaoTbkActivityInfoGetData 结构体
type TaobaoTbkActivityInfoGetData struct {
	// 【本地化】微信推广二维码地址
	WxQrcodeUrl string `json:"wx_qrcode_url,omitempty" xml:"wx_qrcode_url,omitempty"`
	// 淘客推广长链
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 淘客推广短链
	ShortClickUrl string `json:"short_click_url,omitempty" xml:"short_click_url,omitempty"`
	// 投放平台, 1-PC 2-无线
	TerminalType string `json:"terminal_type,omitempty" xml:"terminal_type,omitempty"`
	// 物料素材下载地址
	MaterialOssUrl string `json:"material_oss_url,omitempty" xml:"material_oss_url,omitempty"`
	// 会场名称
	PageName string `json:"page_name,omitempty" xml:"page_name,omitempty"`
	// 活动开始时间
	PageStartTime string `json:"page_start_time,omitempty" xml:"page_start_time,omitempty"`
	// 活动结束时间
	PageEndTime string `json:"page_end_time,omitempty" xml:"page_end_time,omitempty"`
	// 【本地化】微信小程序路径
	WxMiniprogramPath string `json:"wx_miniprogram_path,omitempty" xml:"wx_miniprogram_path,omitempty"`
}

var poolTaobaoTbkActivityInfoGetData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkActivityInfoGetData)
	},
}

// GetTaobaoTbkActivityInfoGetData() 从对象池中获取TaobaoTbkActivityInfoGetData
func GetTaobaoTbkActivityInfoGetData() *TaobaoTbkActivityInfoGetData {
	return poolTaobaoTbkActivityInfoGetData.Get().(*TaobaoTbkActivityInfoGetData)
}

// ReleaseTaobaoTbkActivityInfoGetData 释放TaobaoTbkActivityInfoGetData
func ReleaseTaobaoTbkActivityInfoGetData(v *TaobaoTbkActivityInfoGetData) {
	v.WxQrcodeUrl = ""
	v.ClickUrl = ""
	v.ShortClickUrl = ""
	v.TerminalType = ""
	v.MaterialOssUrl = ""
	v.PageName = ""
	v.PageStartTime = ""
	v.PageEndTime = ""
	v.WxMiniprogramPath = ""
	poolTaobaoTbkActivityInfoGetData.Put(v)
}
