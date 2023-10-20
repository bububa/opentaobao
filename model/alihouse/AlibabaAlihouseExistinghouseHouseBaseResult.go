package alihouse

// AlibabaalihouseexistinghousehousebaseResult 结构体
type AlibabaalihouseexistinghousehousebaseResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果集
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
