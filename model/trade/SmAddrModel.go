package trade

// SmAddrModel 结构体
type SmAddrModel struct {
	// 详细地址，如果地区Code没有填写，API会根据address反向解析地区Code
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 地区Code
	AreaCode string `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 地区名
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 城市Code
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 收货人姓名
	FullName string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// 收货人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 收货人电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 省份名
	PrivinceName string `json:"privince_name,omitempty" xml:"privince_name,omitempty"`
	// 省份Code
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
}
