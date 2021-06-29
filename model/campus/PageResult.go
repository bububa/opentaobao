package campus

// PageResult 
type PageResult struct {
    // content
    Content   *Page `json:"content,omitempty" xml:"content,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误级别
    ErrorLevel   string `json:"error_level,omitempty" xml:"error_level,omitempty"`
    // 请求id
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
