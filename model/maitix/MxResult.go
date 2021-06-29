package maitix

// MxResult 
type MxResult struct {
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 数据对象
    Model   *UnLockTicketResponse `json:"model,omitempty" xml:"model,omitempty"`
    // 结果code，0代表成功，负数代表异常情况
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 是否成功，true:成功，false:失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
