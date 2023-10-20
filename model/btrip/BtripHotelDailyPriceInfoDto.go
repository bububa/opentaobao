package btrip

import (
	"sync"
)

// BtripHotelDailyPriceInfoDto 结构体
type BtripHotelDailyPriceInfoDto struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 餐食信息
	BtripHotelBoardDTO *BtripHotelBoardDto `json:"btrip_hotel_board_d_t_o,omitempty" xml:"btrip_hotel_board_d_t_o,omitempty"`
	// 房价
	CnyPrice int64 `json:"cny_price,omitempty" xml:"cny_price,omitempty"`
	// 不取整的每日优惠后价格
	DiscountDailyPrice int64 `json:"discount_daily_price,omitempty" xml:"discount_daily_price,omitempty"`
	// 取整后的每日优惠后价格
	RoundingDiscountDailyPrice int64 `json:"rounding_discount_daily_price,omitempty" xml:"rounding_discount_daily_price,omitempty"`
}

var poolBtripHotelDailyPriceInfoDto = sync.Pool{
	New: func() any {
		return new(BtripHotelDailyPriceInfoDto)
	},
}

// GetBtripHotelDailyPriceInfoDto() 从对象池中获取BtripHotelDailyPriceInfoDto
func GetBtripHotelDailyPriceInfoDto() *BtripHotelDailyPriceInfoDto {
	return poolBtripHotelDailyPriceInfoDto.Get().(*BtripHotelDailyPriceInfoDto)
}

// ReleaseBtripHotelDailyPriceInfoDto 释放BtripHotelDailyPriceInfoDto
func ReleaseBtripHotelDailyPriceInfoDto(v *BtripHotelDailyPriceInfoDto) {
	v.Date = ""
	v.BtripHotelBoardDTO = nil
	v.CnyPrice = 0
	v.DiscountDailyPrice = 0
	v.RoundingDiscountDailyPrice = 0
	poolBtripHotelDailyPriceInfoDto.Put(v)
}
