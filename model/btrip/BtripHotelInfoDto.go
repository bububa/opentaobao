package btrip

// BtripHotelInfoDto 结构体
type BtripHotelInfoDto struct {
	// 酒店地址
	HotelAddress string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 酒店电话
	HotelTel string `json:"hotel_tel,omitempty" xml:"hotel_tel,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
