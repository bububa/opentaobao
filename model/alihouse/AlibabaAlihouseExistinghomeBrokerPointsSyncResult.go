package alihouse

// AlibabaalihouseexistinghomebrokerpointssyncResult 结构体
type AlibabaalihouseexistinghomebrokerpointssyncResult struct {
	// 操作结果
	Data []BrokerPointResultDto `json:"data,omitempty" xml:"data>broker_point_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
