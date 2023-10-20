package wlb

import (
	"sync"
)

// FreshWaybill 结构体
type FreshWaybill struct {
	// 预计到达时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 简称
	Alias string `json:"alias,omitempty" xml:"alias,omitempty"`
	// 大头笔
	ShortAddress string `json:"short_address,omitempty" xml:"short_address,omitempty"`
	// 交易号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 预留扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 获取的所有电子面单号，以“;”分隔
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
}

var poolFreshWaybill = sync.Pool{
	New: func() any {
		return new(FreshWaybill)
	},
}

// GetFreshWaybill() 从对象池中获取FreshWaybill
func GetFreshWaybill() *FreshWaybill {
	return poolFreshWaybill.Get().(*FreshWaybill)
}

// ReleaseFreshWaybill 释放FreshWaybill
func ReleaseFreshWaybill(v *FreshWaybill) {
	v.Time = ""
	v.Alias = ""
	v.ShortAddress = ""
	v.TradeId = ""
	v.Feature = ""
	v.WaybillCode = ""
	poolFreshWaybill.Put(v)
}
