package mos

// AlibabaMjMoscarnivalReceiveencryptResultDo 
type AlibabaMjMoscarnivalReceiveencryptResultDo struct {

    // 调用链id
    TraceId   string `json:"trace_id,omitempty"`

    // 总行数
    Total   int64 `json:"total,omitempty"`

    // 券结果
    Data   *AlibabaMjMoscarnivalReceiveencryptData `json:"data,omitempty"`

    // 调用是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   int64 `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 结果码
    ResultCode   string `json:"result_code,omitempty"`

}
