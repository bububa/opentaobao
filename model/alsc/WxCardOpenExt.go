package alsc

import (
	"sync"
)

// WxCardOpenExt 结构体
type WxCardOpenExt struct {
	// 等级卡面列表
	WxLevelBgs []WxLevelBgOpenInfo `json:"wx_level_bgs,omitempty" xml:"wx_level_bgs>wx_level_bg_open_info,omitempty"`
	// 品牌logoURL
	BrandLogo string `json:"brand_logo,omitempty" xml:"brand_logo,omitempty"`
	// 统一卡面URL
	GeneralBgLogo string `json:"general_bg_logo,omitempty" xml:"general_bg_logo,omitempty"`
	// 是否统一卡面
	GeneralBgSwitch bool `json:"general_bg_switch,omitempty" xml:"general_bg_switch,omitempty"`
	// 支付后是否可领取
	PaidGetSwitch bool `json:"paid_get_switch,omitempty" xml:"paid_get_switch,omitempty"`
}

var poolWxCardOpenExt = sync.Pool{
	New: func() any {
		return new(WxCardOpenExt)
	},
}

// GetWxCardOpenExt() 从对象池中获取WxCardOpenExt
func GetWxCardOpenExt() *WxCardOpenExt {
	return poolWxCardOpenExt.Get().(*WxCardOpenExt)
}

// ReleaseWxCardOpenExt 释放WxCardOpenExt
func ReleaseWxCardOpenExt(v *WxCardOpenExt) {
	v.WxLevelBgs = v.WxLevelBgs[:0]
	v.BrandLogo = ""
	v.GeneralBgLogo = ""
	v.GeneralBgSwitch = false
	v.PaidGetSwitch = false
	poolWxCardOpenExt.Put(v)
}
