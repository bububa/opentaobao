package alihouse

// AlibabaAlihouseNewhomeProjectQueryResult 结构体
type AlibabaAlihouseNewhomeProjectQueryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *EbbasItemDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
