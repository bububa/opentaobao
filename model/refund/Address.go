package refund

// Address 结构体
type Address struct {
	// 收件人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 邮政编码
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 收件人手机
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 国家
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 省
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 乡镇/街道
	TownName string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	// 乡镇/街道地址详情
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 行政区划代码
	DivisionCode string `json:"division_code,omitempty" xml:"division_code,omitempty"`
	// 地址ID
	AddressId int64 `json:"address_id,omitempty" xml:"address_id,omitempty"`
}
