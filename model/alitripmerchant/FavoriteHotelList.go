package alitripmerchant

// FavoriteHotelList 结构体
type FavoriteHotelList struct {
	// 酒店头图
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// 酒店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店中文名称
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 品牌logo URL
	LogoUrl string `json:"logo_url,omitempty" xml:"logo_url,omitempty"`
	// 卫生健康标识
	MediaUrl string `json:"media_url,omitempty" xml:"media_url,omitempty"`
	// 飞猪标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
