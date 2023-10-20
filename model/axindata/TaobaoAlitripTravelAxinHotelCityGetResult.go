package axindata

// TaobaoalitriptravelaxinhotelcitygetResult 结构体
type TaobaoalitriptravelaxinhotelcitygetResult struct {
	// 城市信息
	DataList []HotelCityVo `json:"data_list,omitempty" xml:"data_list>hotel_city_vo,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
