package alitrippoi

// AlitripPlatformPoiRawFeedResult 结构体
type AlitripPlatformPoiRawFeedResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回值
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// totalRecords
	TotalRecords int64 `json:"total_records,omitempty" xml:"total_records,omitempty"`
	// 成功标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
