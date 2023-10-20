package alihouse

// AlibabaalihousenewhomeprojectadviserbindResult 结构体
type AlibabaalihousenewhomeprojectadviserbindResult struct {
	// 操作结果
	Data []ProjectAdviserBindResultDto `json:"data,omitempty" xml:"data>project_adviser_bind_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
