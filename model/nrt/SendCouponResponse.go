package nrt

// SendCouponResponse 
type SendCouponResponse struct {

    // 面额
    
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    

    // 门槛
    
    StartFee   int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
    

    // 券名称
    
    CouponName   string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
    

    // 券模板ID
    
    CouponTemplateId   int64 `json:"coupon_template_id,omitempty" xml:"coupon_template_id,omitempty"`
    

    // 券实例ID
    
    CouponInstanceId   int64 `json:"coupon_instance_id,omitempty" xml:"coupon_instance_id,omitempty"`
    

    // 结束时间
    
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // 开始时间
    
    StartTime   int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

}
