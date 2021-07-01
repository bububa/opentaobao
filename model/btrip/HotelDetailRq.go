package btrip

// HotelDetailRq 结构体
type HotelDetailRq struct {
	// 入住时间
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店时间
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 城市code
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 标准ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 子渠道ID
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
}
