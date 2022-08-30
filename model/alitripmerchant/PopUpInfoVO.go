package alitripmerchant

// PopUpInfoVO 结构体
type PopUpInfoVO struct {
	// 弹屏类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 抽奖弹屏
	ActivityDrawPopup *ActivityDrawPopupVo `json:"activity_draw_popup,omitempty" xml:"activity_draw_popup,omitempty"`
	// 活动弹窗
	PopUp *PopUpVo `json:"pop_up,omitempty" xml:"pop_up,omitempty"`
	// 优惠券弹屏
	CouponActivityPopUp *CouponActivityPopUpVo `json:"coupon_activity_pop_up,omitempty" xml:"coupon_activity_pop_up,omitempty"`
}
