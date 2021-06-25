package promotion

// CouponTemplateApplyLimitConfig 
type CouponTemplateApplyLimitConfig struct {

    // 优惠券最后能领取时间(超过这个时间，优惠券不能领取)
    ApplyEndTime   string `json:"apply_end_time,omitempty"`

    // 优惠券每日限领数（-1表示不限制）
    CouponDailyLmt   int64 `json:"coupon_daily_lmt,omitempty"`

    // 优惠券总数（-1表示不限制）
    CouponTotalLmt   int64 `json:"coupon_total_lmt,omitempty"`

    // 每人每日限领（-1表示不限制）
    PersonalDailyLmt   int64 `json:"personal_daily_lmt,omitempty"`

    // 个人总领取限制数量（-1表示不限制）
    PersonalLmt   int64 `json:"personal_lmt,omitempty"`

    // 优惠券最早能领取时间(在这时间之前，优惠券不能领取)
    ApplyStartTime   string `json:"apply_start_time,omitempty"`

}
