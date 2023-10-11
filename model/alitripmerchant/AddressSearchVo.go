package alitripmerchant

// AddressSearchVo 结构体
type AddressSearchVo struct {
	// 酒店详情
	HotelList []HotelAddressDetail `json:"hotel_list,omitempty" xml:"hotel_list>hotel_address_detail,omitempty"`
	// 城市列表
	CityList []CityAddressDetail `json:"city_list,omitempty" xml:"city_list>city_address_detail,omitempty"`
}
