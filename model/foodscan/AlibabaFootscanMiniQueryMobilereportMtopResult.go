package foodscan

// AlibabafootscanminiquerymobilereportMtopResult 结构体
type AlibabafootscanminiquerymobilereportMtopResult struct {
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回i数据
	Data *AlibabafootscanminiquerymobilereportData `json:"data,omitempty" xml:"data,omitempty"`
}
