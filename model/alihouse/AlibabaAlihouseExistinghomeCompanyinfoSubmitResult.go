package alihouse

// AlibabaAlihouseExistinghomeCompanyinfoSubmitResult 结构体
type AlibabaAlihouseExistinghomeCompanyinfoSubmitResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 进件id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
