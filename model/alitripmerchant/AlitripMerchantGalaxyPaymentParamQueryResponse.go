package alitripmerchant

// AlitripMerchantGalaxyPaymentParamQueryResponse 
type AlitripMerchantGalaxyPaymentParamQueryResponse struct {
    // 成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 支付参数
    Content   *OrderPayDto `json:"content,omitempty" xml:"content,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
