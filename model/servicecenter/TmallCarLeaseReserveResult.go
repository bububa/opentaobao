package servicecenter

// TmallCarLeaseReserveResult 
type TmallCarLeaseReserveResult struct {

    // 耗时
    CostTime   int64 `json:"cost_time,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误提示
    ErrorMessage   string `json:"error_message,omitempty"`

    // 当前时间
    GmtCurrentTime   int64 `json:"gmt_current_time,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 错误提示
    MsgInfo   string `json:"msg_info,omitempty"`

    // 返回结果，成功或者失败
    Object   bool `json:"object,omitempty"`

    // 返回结果，成功或者失败
    Success   bool `json:"success,omitempty"`

}
