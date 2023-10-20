package alitrippoi

// AlitripplatformpoirawpoioutResult 结构体
type AlitripplatformpoirawpoioutResult struct {
	// 返回素材id
	Datas []AlitripplatformpoirawpoioutData `json:"datas,omitempty" xml:"datas>alitripplatformpoirawpoiout_data,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数(不可用)
	TotalRecords int64 `json:"total_records,omitempty" xml:"total_records,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
