package simba

import (
	"sync"
)

// TimeSpanQueryResVo 结构体
type TimeSpanQueryResVo struct {
	// 折扣时间段
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 折扣
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolTimeSpanQueryResVo = sync.Pool{
	New: func() any {
		return new(TimeSpanQueryResVo)
	},
}

// GetTimeSpanQueryResVo() 从对象池中获取TimeSpanQueryResVo
func GetTimeSpanQueryResVo() *TimeSpanQueryResVo {
	return poolTimeSpanQueryResVo.Get().(*TimeSpanQueryResVo)
}

// ReleaseTimeSpanQueryResVo 释放TimeSpanQueryResVo
func ReleaseTimeSpanQueryResVo(v *TimeSpanQueryResVo) {
	v.Time = ""
	v.Discount = 0
	poolTimeSpanQueryResVo.Put(v)
}
