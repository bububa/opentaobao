package lstmarketing

// AlibabaLstMarketingQuerybyorderidResultDTO 
type AlibabaLstMarketingQuerybyorderidResultDTO struct {
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 订单实体
    Content   *LstTopOrderDTO `json:"content,omitempty" xml:"content,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 执行结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
