package tvpay

// TopResultDO 
type TopResultDO struct {
    // 状态码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 结构体
    Data   *GetLoginInfoByOrderResultDO `json:"data,omitempty" xml:"data,omitempty"`
    // 消息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
