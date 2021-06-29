package logistic

// ResponseResult 
type ResponseResult struct {
    // 返回结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
