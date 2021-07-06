package trade

// ReceiverLogistics 结构体
type ReceiverLogistics struct {
	// 手机号码
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 地区名
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 城市名
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 收货人
	ReceiverFullName string `json:"receiver_full_name,omitempty" xml:"receiver_full_name,omitempty"`
	// 省份名
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// （忽略）
	WholeAddress string `json:"whole_address,omitempty" xml:"whole_address,omitempty"`
	// 收货地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 地区码
	AreaCode int64 `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// divisionId
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
	// 省份码
	ProvinceCode int64 `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 城市码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
}
