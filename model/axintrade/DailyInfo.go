package axintrade

import (
	"sync"
)

// DailyInfo 结构体
type DailyInfo struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 价格，单位为分
	CnyPrice int64 `json:"cny_price,omitempty" xml:"cny_price,omitempty"`
	// 餐食
	BoardDTO *BoardDto `json:"board_d_t_o,omitempty" xml:"board_d_t_o,omitempty"`
	// 原币种价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolDailyInfo = sync.Pool{
	New: func() any {
		return new(DailyInfo)
	},
}

// GetDailyInfo() 从对象池中获取DailyInfo
func GetDailyInfo() *DailyInfo {
	return poolDailyInfo.Get().(*DailyInfo)
}

// ReleaseDailyInfo 释放DailyInfo
func ReleaseDailyInfo(v *DailyInfo) {
	v.Date = ""
	v.CnyPrice = 0
	v.BoardDTO = nil
	v.Price = 0
	poolDailyInfo.Put(v)
}
