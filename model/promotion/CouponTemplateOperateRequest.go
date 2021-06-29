package promotion

// CouponTemplateOperateRequest 
type CouponTemplateOperateRequest struct {
    // 券模版
    CouponTemplate   *CouponTemplate `json:"coupon_template,omitempty" xml:"coupon_template,omitempty"`
    // 用户信息
    UserInfo   *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}
