package promotion

// WdkCouponAbandonParam 
type WdkCouponAbandonParam struct {

    // 券涞源 写死
    
    CouponSource   string `json:"coupon_source,omitempty" xml:"coupon_source,omitempty"`
    

    // 优惠券模版id
    
    TemplateId   string `json:"template_id,omitempty" xml:"template_id,omitempty"`
    

    // 用户id 写死
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 卡券实例id
    
    VoucherId   int64 `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
    

    // 匿名码
    
    MaCode   string `json:"ma_code,omitempty" xml:"ma_code,omitempty"`
    

}
