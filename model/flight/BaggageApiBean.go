package flight

import (
	"sync"
)

// BaggageApiBean 结构体
type BaggageApiBean struct {
	// 单件行李高度限制，单位厘米
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 行李重量，当allWeight为是的时候行李重量描述所有件数的总重量
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 单件行李宽度限制，单位厘米
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 行李总件数，大于等于1的正整数
	Pc int64 `json:"pc,omitempty" xml:"pc,omitempty"`
	// 单件行李长度限制，单位厘米
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 是否所有行李重量
	IsAllWeight bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
}

var poolBaggageApiBean = sync.Pool{
	New: func() any {
		return new(BaggageApiBean)
	},
}

// GetBaggageApiBean() 从对象池中获取BaggageApiBean
func GetBaggageApiBean() *BaggageApiBean {
	return poolBaggageApiBean.Get().(*BaggageApiBean)
}

// ReleaseBaggageApiBean 释放BaggageApiBean
func ReleaseBaggageApiBean(v *BaggageApiBean) {
	v.Height = 0
	v.Weight = 0
	v.Width = 0
	v.Pc = 0
	v.Length = 0
	v.IsAllWeight = false
	poolBaggageApiBean.Put(v)
}
