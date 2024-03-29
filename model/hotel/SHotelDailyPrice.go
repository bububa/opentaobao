package hotel

import (
	"sync"
)

// SHotelDailyPrice 结构体
type SHotelDailyPrice struct {
	// 本shid下所有标准房型的当日库价
	SroomTypeDailyPriceList []SRoomTypeDailyPrice `json:"sroom_type_daily_price_list,omitempty" xml:"sroom_type_daily_price_list>s_room_type_daily_price,omitempty"`
	// 离店日期
	End string `json:"end,omitempty" xml:"end,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 入住时间
	Start string `json:"start,omitempty" xml:"start,omitempty"`
	// 当日所有库价的最低价
	LowPrice int64 `json:"low_price,omitempty" xml:"low_price,omitempty"`
	// 对应标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

var poolSHotelDailyPrice = sync.Pool{
	New: func() any {
		return new(SHotelDailyPrice)
	},
}

// GetSHotelDailyPrice() 从对象池中获取SHotelDailyPrice
func GetSHotelDailyPrice() *SHotelDailyPrice {
	return poolSHotelDailyPrice.Get().(*SHotelDailyPrice)
}

// ReleaseSHotelDailyPrice 释放SHotelDailyPrice
func ReleaseSHotelDailyPrice(v *SHotelDailyPrice) {
	v.SroomTypeDailyPriceList = v.SroomTypeDailyPriceList[:0]
	v.End = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Start = ""
	v.LowPrice = 0
	v.Shid = 0
	poolSHotelDailyPrice.Put(v)
}
