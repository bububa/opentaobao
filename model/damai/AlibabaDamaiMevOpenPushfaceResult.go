package damai

// AlibabaDamaiMevOpenPushfaceResult 
type AlibabaDamaiMevOpenPushfaceResult struct {
    // 错误码
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误内容
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 返回结果
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
