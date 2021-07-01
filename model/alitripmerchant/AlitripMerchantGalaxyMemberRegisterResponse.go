package alitripmerchant

// AlitripMerchantGalaxyMemberRegisterResponse 
type AlitripMerchantGalaxyMemberRegisterResponse struct {
    // 成功还是失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误代码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 是否注册成功
    RegisterResult   bool `json:"register_result,omitempty" xml:"register_result,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
