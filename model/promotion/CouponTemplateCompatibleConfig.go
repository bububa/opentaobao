package promotion

// CouponTemplateCompatibleConfig 结构体
type CouponTemplateCompatibleConfig struct {
	// 是否要求优惠券在一天的23:59:59失效 1表示最后一秒失效
	ValidTillNight int64 `json:"valid_till_night,omitempty" xml:"valid_till_night,omitempty"`
}
