package alitrippoi

// FliggyPoiOutParam 结构体
type FliggyPoiOutParam struct {
	// 开始时间
	StartData string `json:"start_data,omitempty" xml:"start_data,omitempty"`
	// 选择城市名
	CityNames []string `json:"city_names,omitempty" xml:"city_names>string,omitempty"`
	// 城市码可不填
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 开始页数
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
