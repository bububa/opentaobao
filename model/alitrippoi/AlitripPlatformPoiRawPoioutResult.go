package alitrippoi

// AlitripPlatformPoiRawPoioutResult 结构体
type AlitripPlatformPoiRawPoioutResult struct {
	// 总数(不可用)
	TotalRecords int64 `json:"total_records,omitempty" xml:"total_records,omitempty"`
	// 返回素材id
	Datas []AlitripPlatformPoiRawPoioutData `json:"datas,omitempty" xml:"datas>alitrip_platform_poi_raw_poiout_data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
