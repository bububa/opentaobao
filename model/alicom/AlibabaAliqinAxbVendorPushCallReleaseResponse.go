package alicom

// AlibabaaliqinaxbvendorpushcallreleaseResponse 结构体
type AlibabaaliqinaxbvendorpushcallreleaseResponse struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务响应码code为OK时代表请求成功,其他CODE建议做重试机制
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// module，此接口此字段可以不用判断，以外层CODE是否为OK来判断是否调用成
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
