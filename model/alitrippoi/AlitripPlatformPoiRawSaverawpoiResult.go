package alitrippoi

// AlitripplatformpoirawsaverawpoiResult 结构体
type AlitripplatformpoirawsaverawpoiResult struct {
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总记录
	TotalRecords int64 `json:"total_records,omitempty" xml:"total_records,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
