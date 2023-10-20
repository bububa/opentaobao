package alitripmerchant

import (
	"sync"
)

// PrizeDisplay 结构体
type PrizeDisplay struct {
	// 奖品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 奖品URL
	Image string `json:"image,omitempty" xml:"image,omitempty"`
}

var poolPrizeDisplay = sync.Pool{
	New: func() any {
		return new(PrizeDisplay)
	},
}

// GetPrizeDisplay() 从对象池中获取PrizeDisplay
func GetPrizeDisplay() *PrizeDisplay {
	return poolPrizeDisplay.Get().(*PrizeDisplay)
}

// ReleasePrizeDisplay 释放PrizeDisplay
func ReleasePrizeDisplay(v *PrizeDisplay) {
	v.Name = ""
	v.Image = ""
	poolPrizeDisplay.Put(v)
}
