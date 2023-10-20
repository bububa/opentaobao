package tmallservice

import (
	"sync"
)

// SettlementPriceFactor 结构体
type SettlementPriceFactor struct {
	// 计价因子属性
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 计价因子说明
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 计价因子实际值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolSettlementPriceFactor = sync.Pool{
	New: func() any {
		return new(SettlementPriceFactor)
	},
}

// GetSettlementPriceFactor() 从对象池中获取SettlementPriceFactor
func GetSettlementPriceFactor() *SettlementPriceFactor {
	return poolSettlementPriceFactor.Get().(*SettlementPriceFactor)
}

// ReleaseSettlementPriceFactor 释放SettlementPriceFactor
func ReleaseSettlementPriceFactor(v *SettlementPriceFactor) {
	v.Name = ""
	v.Desc = ""
	v.Value = 0
	poolSettlementPriceFactor.Put(v)
}
