package wdk

import (
	"sync"
)

// Condition 结构体
type Condition struct {
	// 满元金额，单位分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 满件门槛
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 是否满件
	CountAt bool `json:"count_at,omitempty" xml:"count_at,omitempty"`
	// 是否指定件数
	CountBegin bool `json:"count_begin,omitempty" xml:"count_begin,omitempty"`
	// 是否第N件
	Appoint bool `json:"appoint,omitempty" xml:"appoint,omitempty"`
	// 是否满元
	AmountAt bool `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
}

var poolCondition = sync.Pool{
	New: func() any {
		return new(Condition)
	},
}

// GetCondition() 从对象池中获取Condition
func GetCondition() *Condition {
	return poolCondition.Get().(*Condition)
}

// ReleaseCondition 释放Condition
func ReleaseCondition(v *Condition) {
	v.Amount = 0
	v.Count = 0
	v.CountAt = false
	v.CountBegin = false
	v.Appoint = false
	v.AmountAt = false
	poolCondition.Put(v)
}
