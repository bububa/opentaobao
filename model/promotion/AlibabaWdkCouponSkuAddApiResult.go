package promotion

// AlibabaWdkCouponSkuAddApiResult 
/* model for simplify = false
type AlibabaWdkCouponSkuAddApiResult struct {

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 返回
    
    Model  *struct {
        CouponTemplateOperateResponse  *CouponTemplateOperateResponse `json:"coupon_template_operate_response,omitempty"`
    } `json:"model,omitempty"`
    

    // 成功标志
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkCouponSkuAddApiResult 
type AlibabaWdkCouponSkuAddApiResult struct {

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 返回
    Model   *CouponTemplateOperateResponse `json:"model,omitempty"`

    // 成功标志
    Success   bool `json:"success,omitempty"`

}
