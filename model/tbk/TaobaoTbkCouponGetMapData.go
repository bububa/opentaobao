package tbk

// TaobaotbkcoupongetMapData 结构体
type TaobaotbkcoupongetMapData struct {
	// 优惠券门槛金额
	Couponstartfee string `json:"coupon_start_fee,omitempty" xml:"coupon_start_fee,omitempty"`
	// 优惠券结束时间
	Couponendtime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 优惠券开始时间
	Couponstarttime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 优惠券金额
	Couponamount string `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
	// 券ID
	Couponactivityid string `json:"coupon_activity_id,omitempty" xml:"coupon_activity_id,omitempty"`
	// 优惠券剩余量
	Couponremaincount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
	// 优惠券总量
	Coupontotalcount int64 `json:"coupon_total_count,omitempty" xml:"coupon_total_count,omitempty"`
	// 券类型，1 表示全网公开券，4 表示妈妈渠道券
	Couponsrcscene int64 `json:"coupon_src_scene,omitempty" xml:"coupon_src_scene,omitempty"`
	// 券属性，0表示店铺券，1表示单品券
	Coupontype int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
}
