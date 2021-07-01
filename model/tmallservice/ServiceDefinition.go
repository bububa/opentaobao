package tmallservice

// ServiceDefinition 结构体
type ServiceDefinition struct {
	// 业务类型
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 服务类型
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
}
