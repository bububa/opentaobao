package btrip

// HotelCityFee 结构体
type HotelCityFee struct {
	// 城市编码，传0代表其他全部城市
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 酒店间夜金额，单位(分)，如需不限请传99999999。注意：差标只能管控到元，角、分会被抹掉，请避免传入角、分的值。
	Fee int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}
