package promotion

// CouponTemplateTimeLimitConfig 
type CouponTemplateTimeLimitConfig struct {
    // 优惠券结束时间
    EndValidTime   string `json:"end_valid_time,omitempty" xml:"end_valid_time,omitempty"`
    // 优惠券开始时间
    StartValidTime   string `json:"start_valid_time,omitempty" xml:"start_valid_time,omitempty"`
    // 优惠券有效时间类型 RANGE(1,"开始/结束时间"), DURATION(2,"固定时长"),
    ValidTimeType   int64 `json:"valid_time_type,omitempty" xml:"valid_time_type,omitempty"`
    // 优惠券有效时长，单位为秒（固定有效时长的优惠券）
    ValidityPeriod   int64 `json:"validity_period,omitempty" xml:"validity_period,omitempty"`
}
