package mos

// AlibabaMjMoscarnivalReceivecouponResultDo 
type AlibabaMjMoscarnivalReceivecouponResultDo struct {
    // 调用链id
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 总行数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // 是否成功
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 结果码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 标题
    Titles   []string `json:"titles,omitempty" xml:"titles>string,omitempty"`
}
