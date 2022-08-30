package tmallcar

// LogisticsTraceReq 结构体
type LogisticsTraceReq struct {
	// 当前区县
	CurrentDistinct string `json:"current_distinct,omitempty" xml:"current_distinct,omitempty"`
	// 当前详细地址
	CurrentLocationDetail string `json:"current_location_detail,omitempty" xml:"current_location_detail,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 当前城市
	CurrentCity string `json:"current_city,omitempty" xml:"current_city,omitempty"`
	// 当街街道
	CurrentStreet string `json:"current_street,omitempty" xml:"current_street,omitempty"`
	// 当前省份
	CurrentProvince string `json:"current_province,omitempty" xml:"current_province,omitempty"`
	// 物流轨迹详细描述
	LogisticsDesc string `json:"logistics_desc,omitempty" xml:"logistics_desc,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}
