package xhotelitem

import (
	"sync"
)

// Rateinfos 结构体
type Rateinfos struct {
	// 外部房源id
	OutRid string `json:"out_rid,omitempty" xml:"out_rid,omitempty"`
	// 外部rpcode
	RatePlanCode string `json:"rate_plan_code,omitempty" xml:"rate_plan_code,omitempty"`
}

var poolRateinfos = sync.Pool{
	New: func() any {
		return new(Rateinfos)
	},
}

// GetRateinfos() 从对象池中获取Rateinfos
func GetRateinfos() *Rateinfos {
	return poolRateinfos.Get().(*Rateinfos)
}

// ReleaseRateinfos 释放Rateinfos
func ReleaseRateinfos(v *Rateinfos) {
	v.OutRid = ""
	v.RatePlanCode = ""
	poolRateinfos.Put(v)
}
