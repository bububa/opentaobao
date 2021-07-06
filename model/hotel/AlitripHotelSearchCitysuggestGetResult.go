package hotel

// AlitripHotelSearchCitysuggestGetResult 结构体
type AlitripHotelSearchCitysuggestGetResult struct {
	// mappingCode
	MappingCode string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误消息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// bizExtMap
	BizExtMap *Bizextmap `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// headers
	Headers *Headers `json:"headers,omitempty" xml:"headers,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 具体的返回数据
	Model *TopCitySuggestListVo `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
