package damai

// MpResult 
type MpResult struct {
    // model
    Model   *QueryProjectResult `json:"model,omitempty" xml:"model,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
