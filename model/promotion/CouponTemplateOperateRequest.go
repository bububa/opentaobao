package promotion

// CouponTemplateOperateRequest 
/* model for simplify = false
type CouponTemplateOperateRequest struct {

    // 券模版
    
    CouponTemplate  *struct {
        CouponTemplate  *CouponTemplate `json:"coupon_template,omitempty"`
    } `json:"coupon_template,omitempty"`
    

    // 用户信息
    
    UserInfo  *struct {
        UserInfo  *UserInfo `json:"user_info,omitempty"`
    } `json:"user_info,omitempty"`
    

}
*/

// CouponTemplateOperateRequest 
type CouponTemplateOperateRequest struct {

    // 券模版
    CouponTemplate   *CouponTemplate `json:"coupon_template,omitempty"`

    // 用户信息
    UserInfo   *UserInfo `json:"user_info,omitempty"`

}
