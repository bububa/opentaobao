package tbk

// TaobaoTbkCouponGetMapData 
type TaobaoTbkCouponGetMapData struct {

    // 优惠券门槛金额
    
    CouponStartFee   string `json:"coupon_start_fee,omitempty" xml:"coupon_start_fee,omitempty"`
    

    // 优惠券剩余量
    
    CouponRemainCount   int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
    

    // 优惠券总量
    
    CouponTotalCount   int64 `json:"coupon_total_count,omitempty" xml:"coupon_total_count,omitempty"`
    

    // 优惠券结束时间
    
    CouponEndTime   string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
    

    // 优惠券开始时间
    
    CouponStartTime   string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
    

    // 优惠券金额
    
    CouponAmount   string `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
    

    // 券类型，1 表示全网公开券，4 表示妈妈渠道券
    
    CouponSrcScene   int64 `json:"coupon_src_scene,omitempty" xml:"coupon_src_scene,omitempty"`
    

    // 券属性，0表示店铺券，1表示单品券
    
    CouponType   int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
    

    // 券ID
    
    CouponActivityId   string `json:"coupon_activity_id,omitempty" xml:"coupon_activity_id,omitempty"`
    

}
