package aliospay

// AliOSPayResponse 
type AliOSPayResponse struct {
    // 请求唯一id，不可重复，服务端会根据此参数防重放
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 错误码
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 业务数据
    Data   *SearchRecordResponse `json:"data,omitempty" xml:"data,omitempty"`
}
