package alitripmerchant

// HotelListSearchDto 结构体
type HotelListSearchDto struct {
	// 酒店设施
	HotelFacilityList []string `json:"hotel_facility_list,omitempty" xml:"hotel_facility_list>string,omitempty"`
	// 娱乐设施
	FunFacilityList []string `json:"fun_facility_list,omitempty" xml:"fun_facility_list>string,omitempty"`
	// 酒店服务
	HotelServiceList []string `json:"hotel_service_list,omitempty" xml:"hotel_service_list>string,omitempty"`
	// 设施icon
	FacilityList []FacilityDto `json:"facility_list,omitempty" xml:"facility_list>facility_dto,omitempty"`
	// 品牌名字
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 地址详情
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店星级
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 图片头图
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// 酒店中文名
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 国内外 0国内 1国外
	Domestic string `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 品牌头图
	BrandUrl string `json:"brand_url,omitempty" xml:"brand_url,omitempty"`
	// 城市中文名
	CityCn string `json:"city_cn,omitempty" xml:"city_cn,omitempty"`
	// 国家中文名
	CountryCn string `json:"country_cn,omitempty" xml:"country_cn,omitempty"`
	// 最低价格房型名
	RoomNameCn string `json:"room_name_cn,omitempty" xml:"room_name_cn,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 品牌code
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 经纬度类型 0-高德 1-google
	PositionType int64 `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 标准库id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 酒店商品的最低价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 酒店id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 是否热门
	Hot int64 `json:"hot,omitempty" xml:"hot,omitempty"`
	// 是否满房 0未满房1满房
	Full bool `json:"full,omitempty" xml:"full,omitempty"`
}
