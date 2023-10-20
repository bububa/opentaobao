package axintrade

import (
	"sync"
)

// DailyPriceInfoDto 结构体
type DailyPriceInfoDto struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 餐食
	Board *Board `json:"board,omitempty" xml:"board,omitempty"`
	// 房间价格（人民币元）
	CnyPrice int64 `json:"cny_price,omitempty" xml:"cny_price,omitempty"`
	// 原币种金额
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolDailyPriceInfoDto = sync.Pool{
	New: func() any {
		return new(DailyPriceInfoDto)
	},
}

// GetDailyPriceInfoDto() 从对象池中获取DailyPriceInfoDto
func GetDailyPriceInfoDto() *DailyPriceInfoDto {
	return poolDailyPriceInfoDto.Get().(*DailyPriceInfoDto)
}

// ReleaseDailyPriceInfoDto 释放DailyPriceInfoDto
func ReleaseDailyPriceInfoDto(v *DailyPriceInfoDto) {
	v.Date = ""
	v.Board = nil
	v.CnyPrice = 0
	v.Price = 0
	poolDailyPriceInfoDto.Put(v)
}
