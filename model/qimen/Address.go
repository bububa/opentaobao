package qimen

// Address 结构体
type Address struct {
	// 区域
	Region string `json:"region,omitempty" xml:"region,omitempty"`
	// 区域
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 国家二字码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 村镇
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
}
