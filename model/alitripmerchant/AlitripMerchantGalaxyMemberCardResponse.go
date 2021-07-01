package alitripmerchant

// AlitripMerchantGalaxyMemberCardResponse 
type AlitripMerchantGalaxyMemberCardResponse struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 返回参数
    Content   *MemberCardDto `json:"content,omitempty" xml:"content,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
