package promotion

// AlibabaWdkCouponTemplateCreateApiResult 
type AlibabaWdkCouponTemplateCreateApiResult struct {

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 返回
    Model   *CouponTemplateOperateResponse `json:"model,omitempty"`

    // 成功标志
    Success   bool `json:"success,omitempty"`

}
