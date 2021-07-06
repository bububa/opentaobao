package hotel

// AlitripHotelDetailInfoGetResult 结构体
type AlitripHotelDetailInfoGetResult struct {
	// mappingCode
	MappingCode string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// bizExtMap
	BizExtMap *Bizextmap `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// headers
	Headers *Headers `json:"headers,omitempty" xml:"headers,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 返回结果
	Model *HotelTopDetailsVo `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
