package wdk

// CouponStatisticsResultDO 
/* model for simplify = false
type CouponStatisticsResultDO struct {

    // 核券量
    
    UseCouponCount   int64 `json:"use_coupon_count,omitempty"`
    

    // 发券量
    
    SendCouponCount   int64 `json:"send_coupon_count,omitempty"`
    

    // 券id
    
    CouponId   string `json:"coupon_id,omitempty"`
    

    // 券名称
    
    CouponName   string `json:"coupon_name,omitempty"`
    

    // 导购员id
    
    GuideId   string `json:"guide_id,omitempty"`
    

    // 日期
    
    StatisticsDate   string `json:"statistics_date,omitempty"`
    

    // 业务标识
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

}
*/

// CouponStatisticsResultDO 
type CouponStatisticsResultDO struct {

    // 核券量
    UseCouponCount   int64 `json:"use_coupon_count,omitempty"`

    // 发券量
    SendCouponCount   int64 `json:"send_coupon_count,omitempty"`

    // 券id
    CouponId   string `json:"coupon_id,omitempty"`

    // 券名称
    CouponName   string `json:"coupon_name,omitempty"`

    // 导购员id
    GuideId   string `json:"guide_id,omitempty"`

    // 日期
    StatisticsDate   string `json:"statistics_date,omitempty"`

    // 业务标识
    MerchantCode   string `json:"merchant_code,omitempty"`

}
