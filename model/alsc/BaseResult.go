package alsc

// BaseResult 
type BaseResult struct {

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否失败
    Fail   bool `json:"fail,omitempty"`

    // 是否处理中
    Process   bool `json:"process,omitempty"`

    // 1成功，2失败，3处理中
    Status   int64 `json:"status,omitempty"`

    // 是否执行成功
    Success   bool `json:"success,omitempty"`

}
