package axindata

// StdHotelDto 结构体
type StdHotelDto struct {
	// 标准酒店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 酒店英文名称
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 英文地址
	AddressEn string `json:"address_en,omitempty" xml:"address_en,omitempty"`
	// 经度
	Longtitude string `json:"longtitude,omitempty" xml:"longtitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 酒店档次0-暂无，1-经济连锁，2-二星及以下，3-舒适，4-高档，5-豪华
	Star string `json:"star,omitempty" xml:"star,omitempty"`
	// 开业时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 装修时间
	DecorateTime string `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	// 酒店描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 英文描述
	DescriptionEn string `json:"description_en,omitempty" xml:"description_en,omitempty"`
	// 电话
	HotelTel string `json:"hotel_tel,omitempty" xml:"hotel_tel,omitempty"`
	// 标准酒店ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 0-高德（国内酒店） 1-google(国际酒店)，默认0
	PositionType int64 `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 0--酒店，1--客栈，2--非标民宿
	HotelType int64 `json:"hotel_type,omitempty" xml:"hotel_type,omitempty"`
	// 国别 0-国内 1-国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	//  0,营业中；-1，筹建中；-2，暂停营业；-3，已停业；默认为0
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 城市信息
	City *DivisionDto `json:"city,omitempty" xml:"city,omitempty"`
	// 省信息
	Province *DivisionDto `json:"province,omitempty" xml:"province,omitempty"`
	// 国家信息
	Country *DivisionDto `json:"country,omitempty" xml:"country,omitempty"`
}
