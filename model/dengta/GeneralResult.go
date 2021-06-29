package dengta

// GeneralResult 
type GeneralResult struct {
    // 请求标识
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 返回码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 是否成功
    Value   bool `json:"value,omitempty" xml:"value,omitempty"`
}
