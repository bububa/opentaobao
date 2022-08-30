package waybill

// AddressDto 结构体
type AddressDto struct {
	// 镇/街道
	TownName string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 省
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 城市，长度小于20
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 详细地址，长度小于256
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 区，长度小于20
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 省，长度小于20
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道，长度小于30
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 订购关系id
	WaybillAddressId string `json:"waybill_address_id,omitempty" xml:"waybill_address_id,omitempty"`
}
