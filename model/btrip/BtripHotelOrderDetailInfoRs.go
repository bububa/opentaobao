package btrip

// BtripHotelOrderDetailInfoRs 结构体
type BtripHotelOrderDetailInfoRs struct {
	// 每日房价信息
	DailyPriceInfoList []BtripHotelDailyPriceInfoDto `json:"daily_price_info_list,omitempty" xml:"daily_price_info_list>btrip_hotel_daily_price_info_dto,omitempty"`
	// 订单酒店信息
	BtripHotelInfo *BtripHotelInfoDto `json:"btrip_hotel_info,omitempty" xml:"btrip_hotel_info,omitempty"`
	// 订单主体信息
	BtripHotelOrderMainInfo *BtripHotelOrderMainInfoDto `json:"btrip_hotel_order_main_info,omitempty" xml:"btrip_hotel_order_main_info,omitempty"`
	// 房型信息
	BtripHotelRoomInfo *BtripHotelRoomInfoDto `json:"btrip_hotel_room_info,omitempty" xml:"btrip_hotel_room_info,omitempty"`
}
