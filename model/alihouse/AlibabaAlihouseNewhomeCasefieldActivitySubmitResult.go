package alihouse

// AlibabaAlihouseNewhomeCasefieldActivitySubmitResult 结构体
type AlibabaAlihouseNewhomeCasefieldActivitySubmitResult struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回对象
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
