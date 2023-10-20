package alihouse

// AlibabaalihousemerchantopenupdateResult 结构体
type AlibabaalihousemerchantopenupdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
