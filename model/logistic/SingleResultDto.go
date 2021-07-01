package logistic

// SingleResultDto 
type SingleResultDto struct {
    // 是否需要重试
    IsRetry   bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
    // 业务返回结果
    Result   *ExpressModifyAppointTopResponseDto `json:"result,omitempty" xml:"result,omitempty"`
    // 错误描述
    ErrorDesc   string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
    // 调用码
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 错误编码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 是否幂等
    IsIdempotent   bool `json:"is_idempotent,omitempty" xml:"is_idempotent,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
