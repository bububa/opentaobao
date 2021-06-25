package servicecenter

// TmallCarLeaseItemcarinfoResult 
type TmallCarLeaseItemcarinfoResult struct {

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

    // 返回的对象
    Object   *CarItemInfoDto `json:"object,omitempty"`

    // 成功与否
    Success   bool `json:"success,omitempty"`

}
