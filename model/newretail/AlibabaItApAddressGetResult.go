package newretail

// AlibabaItApAddressGetResult 
type AlibabaItApAddressGetResult struct {
    // 返回的位置结构体
    Data   *ApAddressInfo `json:"data,omitempty" xml:"data,omitempty"`
    // 返回的错误码
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 返回的错误message
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 接口返回成功与否
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
