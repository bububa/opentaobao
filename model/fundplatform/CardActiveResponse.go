package fundplatform

// CardActiveResponse 
type CardActiveResponse struct {
    // 错误码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 错误消息
    ResultMessage   string `json:"result_message,omitempty" xml:"result_message,omitempty"`
    // 是否处理成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
