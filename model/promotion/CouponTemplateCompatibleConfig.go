package promotion

// CouponTemplateCompatibleConfig 
/* model for simplify = false
type CouponTemplateCompatibleConfig struct {

    // 是否要求优惠券在一天的23:59:59失效 1表示最后一秒失效
    
    ValidTillNight   int64 `json:"valid_till_night,omitempty"`
    

}
*/

// CouponTemplateCompatibleConfig 
type CouponTemplateCompatibleConfig struct {

    // 是否要求优惠券在一天的23:59:59失效 1表示最后一秒失效
    ValidTillNight   int64 `json:"valid_till_night,omitempty"`

}
