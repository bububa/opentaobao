package alitripmerchant

// AddressListSearchDto 
type AddressListSearchDto struct {
    // 城市列表
    CityList   []AddressLetterDto `json:"city_list,omitempty" xml:"city_list>address_letter_dto,omitempty"`
    // 热门城市
    HotCityList   []AddressSearchDto `json:"hot_city_list,omitempty" xml:"hot_city_list>address_search_dto,omitempty"`
}
