package promotion

// Activity 结构体
type Activity struct {
	// enabled代表有效，invalid代表失效。other代表空值
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 领用优惠券的链接
	ActivityUrl string `json:"activity_url,omitempty" xml:"activity_url,omitempty"`
	// self代表自己创建，other他人创建
	CreateUser string `json:"create_user,omitempty" xml:"create_user,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动对应的优惠券ID
	CouponId int64 `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 卖家设置优惠券领取的总领用量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 已经领取的优惠券的数量
	AppliedCount int64 `json:"applied_count,omitempty" xml:"applied_count,omitempty"`
	// 每个买家限领取优惠券的数量，1～5张
	PersonLimitCount int64 `json:"person_limit_count,omitempty" xml:"person_limit_count,omitempty"`
}
