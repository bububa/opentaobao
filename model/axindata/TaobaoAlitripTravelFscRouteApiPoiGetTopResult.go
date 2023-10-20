package axindata

// TaobaoalitriptravelfscrouteapipoigetTopResult 结构体
type TaobaoalitriptravelfscrouteapipoigetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 业务数据
	FscPoiApiResponse *FscPoiApiResponse `json:"fsc_poi_api_response,omitempty" xml:"fsc_poi_api_response,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
