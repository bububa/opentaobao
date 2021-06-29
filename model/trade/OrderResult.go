package trade

// OrderResult 
type OrderResult struct {
    // 错误编码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误消息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否取消成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 取消的订单
    Result   *OrderObject `json:"result,omitempty" xml:"result,omitempty"`
    // 创建的订单
    Trade   *TradeOrder `json:"trade,omitempty" xml:"trade,omitempty"`
}
