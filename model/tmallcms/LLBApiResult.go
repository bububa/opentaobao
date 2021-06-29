package tmallcms

// LLBApiResult 
type LLBApiResult struct {
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 结果对象
    Model   *SpreadLinkDO `json:"model,omitempty" xml:"model,omitempty"`
    // 错误码
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 成功结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
