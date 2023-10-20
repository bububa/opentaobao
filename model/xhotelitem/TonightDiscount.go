package xhotelitem

import (
	"sync"
)

// TonightDiscount 结构体
type TonightDiscount struct {
	// 活动折扣
	InvestmentNumber string `json:"investment_number,omitempty" xml:"investment_number,omitempty"`
	// 起始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

var poolTonightDiscount = sync.Pool{
	New: func() any {
		return new(TonightDiscount)
	},
}

// GetTonightDiscount() 从对象池中获取TonightDiscount
func GetTonightDiscount() *TonightDiscount {
	return poolTonightDiscount.Get().(*TonightDiscount)
}

// ReleaseTonightDiscount 释放TonightDiscount
func ReleaseTonightDiscount(v *TonightDiscount) {
	v.InvestmentNumber = ""
	v.StartTime = ""
	poolTonightDiscount.Put(v)
}
