package btrip

// BtripHotelDailyPriceInfoDto 结构体
type BtripHotelDailyPriceInfoDto struct {
	// 餐食信息
	BtripHotelBoardDTO *BtripHotelBoardDto `json:"btrip_hotel_board_d_t_o,omitempty" xml:"btrip_hotel_board_d_t_o,omitempty"`
	// 房间价格,人民币，单位分
	CnyPrice int64 `json:"cny_price,omitempty" xml:"cny_price,omitempty"`
	// 入住日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
}
