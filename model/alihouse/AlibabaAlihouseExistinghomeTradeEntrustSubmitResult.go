package alihouse

// AlibabaAlihouseExistinghomeTradeEntrustSubmitResult 结构体
type AlibabaAlihouseExistinghomeTradeEntrustSubmitResult struct {
	// 消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回值
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
