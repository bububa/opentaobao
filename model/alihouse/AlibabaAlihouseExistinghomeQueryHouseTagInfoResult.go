package alihouse

// AlibabaalihouseexistinghomequeryhousetaginfoResult 结构体
type AlibabaalihouseexistinghomequeryhousetaginfoResult struct {
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回值
	Data *SyncHouseTagFeatureDto `json:"data,omitempty" xml:"data,omitempty"`
	// 1
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
