package axindata

// HotelPriceVo 结构体
type HotelPriceVo struct {
	// 报价信息
	RoomList []RoomPriceVo `json:"room_list,omitempty" xml:"room_list>room_price_vo,omitempty"`
	// 入住日期
	CheckIn string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// 离店日期
	CheckOut string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
