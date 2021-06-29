package ieagency

// BaseApiResult 
type BaseApiResult struct {
    // 错误编码
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 创建支付订单结果
    Model   *IeBookPayOrderVo `json:"model,omitempty" xml:"model,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
