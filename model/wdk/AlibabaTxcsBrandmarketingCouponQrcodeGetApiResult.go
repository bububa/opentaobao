package wdk

// AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult 
/* model for simplify = false
type AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult struct {

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 返回内容
    
    Model  *struct {
        CouponQrcodeResultDO  *CouponQrcodeResultDO `json:"coupon_qrcode_result_do,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult 
type AlibabaTxcsBrandmarketingCouponQrcodeGetApiResult struct {

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 返回内容
    Model   *CouponQrcodeResultDO `json:"model,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
