package btrip

// HotHotelIdListRs 结构体
type HotHotelIdListRs struct {
	// 酒店Id列表
	HotelIds []string `json:"hotel_ids,omitempty" xml:"hotel_ids>string,omitempty"`
}
