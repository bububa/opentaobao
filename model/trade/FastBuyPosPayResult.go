package trade

// FastBuyPosPayResult 结构体
type FastBuyPosPayResult struct {
	// 返回的错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 错误描述
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
