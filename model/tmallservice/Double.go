package tmallservice

import (
	"sync"
)

// Double 结构体
type Double struct {
	// 费用名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 费用金额（分）
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolDouble = sync.Pool{
	New: func() any {
		return new(Double)
	},
}

// GetDouble() 从对象池中获取Double
func GetDouble() *Double {
	return poolDouble.Get().(*Double)
}

// ReleaseDouble 释放Double
func ReleaseDouble(v *Double) {
	v.Name = ""
	v.Amount = 0
	poolDouble.Put(v)
}
