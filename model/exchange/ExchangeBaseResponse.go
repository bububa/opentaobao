package exchange

// ExchangeBaseResponse 
type ExchangeBaseResponse struct {
    // 返回结果说明
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 返回结果码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 是否成功调用
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 换货单信息
    Exchange   *Exchange `json:"exchange,omitempty" xml:"exchange,omitempty"`
}