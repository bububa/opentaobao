package alitripmerchant

import (
	"sync"
)

// PopUpInfo 结构体
type PopUpInfo struct {
	// 弹屏图片
	PopUpUrl string `json:"pop_up_url,omitempty" xml:"pop_up_url,omitempty"`
	// 弹屏跳转链接
	PopUpRedirectUrl string `json:"pop_up_redirect_url,omitempty" xml:"pop_up_redirect_url,omitempty"`
}

var poolPopUpInfo = sync.Pool{
	New: func() any {
		return new(PopUpInfo)
	},
}

// GetPopUpInfo() 从对象池中获取PopUpInfo
func GetPopUpInfo() *PopUpInfo {
	return poolPopUpInfo.Get().(*PopUpInfo)
}

// ReleasePopUpInfo 释放PopUpInfo
func ReleasePopUpInfo(v *PopUpInfo) {
	v.PopUpUrl = ""
	v.PopUpRedirectUrl = ""
	poolPopUpInfo.Put(v)
}
