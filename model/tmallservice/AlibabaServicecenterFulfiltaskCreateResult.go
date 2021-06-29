package tmallservice

// AlibabaServicecenterFulfiltaskCreateResult 
type AlibabaServicecenterFulfiltaskCreateResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 核销单id
    FulfilTaskId   int64 `json:"fulfil_task_id,omitempty" xml:"fulfil_task_id,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
