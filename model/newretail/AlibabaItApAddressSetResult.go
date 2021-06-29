package newretail

// AlibabaItApAddressSetResult 
type AlibabaItApAddressSetResult struct {
    // 返回内容
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 返回结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 返回的错误码
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 返回的message
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
