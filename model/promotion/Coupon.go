package promotion

// Coupon 
type Coupon struct {
    // 优惠券ID
    CouponId   int64 `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
    // 优惠券的面值，返回的是“分”，不是“元”，500代表500分相当于5元
    Denominations   int64 `json:"denominations,omitempty" xml:"denominations,omitempty"`
    // 优惠券创建时间
    CreatTime   string `json:"creat_time,omitempty" xml:"creat_time,omitempty"`
    // 优惠券的截止日期
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 订单满多少分才能用这个优惠券，501就是满501分能使用。注意：返回的是“分”，不是“元”
    Condition   int64 `json:"condition,omitempty" xml:"condition,omitempty"`
    // 优惠券的创建渠道，自己创建/他人创建
    CreateChannel   string `json:"create_channel,omitempty" xml:"create_channel,omitempty"`
}
