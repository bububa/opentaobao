package middleclaims

// DeliveryAddressDto 结构体
type DeliveryAddressDto struct {
	// 收货人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收货人手机号码
	ReceiverMobilePhone string `json:"receiver_mobile_phone,omitempty" xml:"receiver_mobile_phone,omitempty"`
	// 国家
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 省份
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 城市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区域
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 邮政编码
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 具体地址
	ReceiverAddressDetail string `json:"receiver_address_detail,omitempty" xml:"receiver_address_detail,omitempty"`
}
