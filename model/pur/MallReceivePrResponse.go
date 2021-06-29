package pur

// MallReceivePrResponse 
type MallReceivePrResponse struct {
    // 返回数据
    Data   *MallReceivePrResponseData `json:"data,omitempty" xml:"data,omitempty"`
    // 错误代码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误提示
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 返回标识
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
