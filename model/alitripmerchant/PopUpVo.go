package alitripmerchant

import (
	"sync"
)

// PopUpVo 结构体
type PopUpVo struct {
	// 弹屏跳转链接
	PopUpRedirectUrl string `json:"pop_up_redirect_url,omitempty" xml:"pop_up_redirect_url,omitempty"`
	// 弹屏图片
	PopUpUrl string `json:"pop_up_url,omitempty" xml:"pop_up_url,omitempty"`
	// 活动id
	OfferId int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
}

var poolPopUpVo = sync.Pool{
	New: func() any {
		return new(PopUpVo)
	},
}

// GetPopUpVo() 从对象池中获取PopUpVo
func GetPopUpVo() *PopUpVo {
	return poolPopUpVo.Get().(*PopUpVo)
}

// ReleasePopUpVo 释放PopUpVo
func ReleasePopUpVo(v *PopUpVo) {
	v.PopUpRedirectUrl = ""
	v.PopUpUrl = ""
	v.OfferId = 0
	poolPopUpVo.Put(v)
}
