package alihouse

// AlibabaAlihouseNewhomeProjectSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectSubmitResult struct {
	// 接口返回对象
	Data *EbbasProjectSubmitVO `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
