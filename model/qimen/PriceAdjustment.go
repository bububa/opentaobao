package qimen

import (
	"sync"
)

// PriceAdjustment 结构体
type PriceAdjustment struct {
	// test
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// test
	StandardPrice string `json:"standardPrice,omitempty" xml:"standardPrice,omitempty"`
	// test
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// test
	StartDate string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// test
	EndDate string `json:"endDate,omitempty" xml:"endDate,omitempty"`
}

var poolPriceAdjustment = sync.Pool{
	New: func() any {
		return new(PriceAdjustment)
	},
}

// GetPriceAdjustment() 从对象池中获取PriceAdjustment
func GetPriceAdjustment() *PriceAdjustment {
	return poolPriceAdjustment.Get().(*PriceAdjustment)
}

// ReleasePriceAdjustment 释放PriceAdjustment
func ReleasePriceAdjustment(v *PriceAdjustment) {
	v.Type = ""
	v.StandardPrice = ""
	v.Discount = ""
	v.StartDate = ""
	v.EndDate = ""
	poolPriceAdjustment.Put(v)
}
