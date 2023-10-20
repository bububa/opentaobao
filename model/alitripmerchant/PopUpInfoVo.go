package alitripmerchant

import (
	"sync"
)

// PopUpInfoVo 结构体
type PopUpInfoVo struct {
	// 弹屏类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 抽奖弹屏
	ActivityDrawPopup *ActivityDrawPopupVo `json:"activity_draw_popup,omitempty" xml:"activity_draw_popup,omitempty"`
	// 活动弹窗
	PopUp *PopUpVo `json:"pop_up,omitempty" xml:"pop_up,omitempty"`
	// 优惠券弹屏
	CouponActivityPopUp *CouponActivityPopUpVo `json:"coupon_activity_pop_up,omitempty" xml:"coupon_activity_pop_up,omitempty"`
}

var poolPopUpInfoVo = sync.Pool{
	New: func() any {
		return new(PopUpInfoVo)
	},
}

// GetPopUpInfoVo() 从对象池中获取PopUpInfoVo
func GetPopUpInfoVo() *PopUpInfoVo {
	return poolPopUpInfoVo.Get().(*PopUpInfoVo)
}

// ReleasePopUpInfoVo 释放PopUpInfoVo
func ReleasePopUpInfoVo(v *PopUpInfoVo) {
	v.Type = ""
	v.ActivityDrawPopup = nil
	v.PopUp = nil
	v.CouponActivityPopUp = nil
	poolPopUpInfoVo.Put(v)
}
