package logistic

// ResultDto 结构体
type ResultDto struct {
	// 快递公司资源列表
	ResourceList []LogisticsResourceDto `json:"resource_list,omitempty" xml:"resource_list>logistics_resource_dto,omitempty"`
	// 错误信息
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
