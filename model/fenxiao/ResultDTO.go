package fenxiao

// ResultDTO 
type ResultDTO struct {
    // 执行结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 下架结果
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
    // 错误码
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 链路ID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 异常名
    ExpName   string `json:"exp_name,omitempty" xml:"exp_name,omitempty"`
    // 重定向url
    RedirectUrl   string `json:"redirect_url,omitempty" xml:"redirect_url,omitempty"`
}
