package alihouse

// AlibabaAlihouseExistinghomePosOpenSubmitResult 结构体
type AlibabaAlihouseExistinghomePosOpenSubmitResult struct {
	// 返回编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回值
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
