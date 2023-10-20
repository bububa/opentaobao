package alitripmerchant

import (
	"sync"
)

// CouponActivityPopUpVo 结构体
type CouponActivityPopUpVo struct {
	// 优惠券模板
	CouponTemplateVOList []DisplayCouponTemplateVo `json:"coupon_template_v_o_list,omitempty" xml:"coupon_template_v_o_list>display_coupon_template_vo,omitempty"`
	// 1
	HomePagePopup string `json:"home_page_popup,omitempty" xml:"home_page_popup,omitempty"`
	// 1
	DetailsPagePicture string `json:"details_page_picture,omitempty" xml:"details_page_picture,omitempty"`
	// 活动id
	PmActivityId int64 `json:"pm_activity_id,omitempty" xml:"pm_activity_id,omitempty"`
}

var poolCouponActivityPopUpVo = sync.Pool{
	New: func() any {
		return new(CouponActivityPopUpVo)
	},
}

// GetCouponActivityPopUpVo() 从对象池中获取CouponActivityPopUpVo
func GetCouponActivityPopUpVo() *CouponActivityPopUpVo {
	return poolCouponActivityPopUpVo.Get().(*CouponActivityPopUpVo)
}

// ReleaseCouponActivityPopUpVo 释放CouponActivityPopUpVo
func ReleaseCouponActivityPopUpVo(v *CouponActivityPopUpVo) {
	v.CouponTemplateVOList = v.CouponTemplateVOList[:0]
	v.HomePagePopup = ""
	v.DetailsPagePicture = ""
	v.PmActivityId = 0
	poolCouponActivityPopUpVo.Put(v)
}
