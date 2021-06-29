package wdk

// AlibabaWdkCouponTemplateQueryumpactidApiResult 
type AlibabaWdkCouponTemplateQueryumpactidApiResult struct {
    // 数据
    Models   []CouponRelatedResponse `json:"models,omitempty" xml:"models>coupon_related_response,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // true为成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
