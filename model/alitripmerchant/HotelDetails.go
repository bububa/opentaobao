package alitripmerchant

// HotelDetails 结构体
type HotelDetails struct {
	// 酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 酒店头图url
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// address
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 酒店中文名称
	NameCn string `json:"name_cn,omitempty" xml:"name_cn,omitempty"`
	// 标准库id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
