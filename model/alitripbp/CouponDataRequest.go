package alitripbp

// CouponDataRequest 结构体
type CouponDataRequest struct {
	// 券id
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 券额度(单位:分)
	CouponPrice string `json:"coupon_price,omitempty" xml:"coupon_price,omitempty"`
	// 券失效时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 场景
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 券开始有效时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 外部用户唯一标识
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 券模版标识
	CouponTemplateId string `json:"coupon_template_id,omitempty" xml:"coupon_template_id,omitempty"`
	// 额外信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 券状态
	CouponStatus int64 `json:"coupon_status,omitempty" xml:"coupon_status,omitempty"`
}
