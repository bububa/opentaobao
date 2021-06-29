package cityretail

// WorkResult 
type WorkResult struct {
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 请求结果数据
    Data   *ChangeOrderResponseDto `json:"data,omitempty" xml:"data,omitempty"`
    // 标示服务成功/失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 返回数据对象
    ResultData   *OrderLogisticsResponseDto `json:"result_data,omitempty" xml:"result_data,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}
