package fenxiao

import (
	"sync"
)

// TopProductStatusResult 结构体
type TopProductStatusResult struct {
	// 产品铺货状态
	Release bool `json:"release,omitempty" xml:"release,omitempty"`
}

var poolTopProductStatusResult = sync.Pool{
	New: func() any {
		return new(TopProductStatusResult)
	},
}

// GetTopProductStatusResult() 从对象池中获取TopProductStatusResult
func GetTopProductStatusResult() *TopProductStatusResult {
	return poolTopProductStatusResult.Get().(*TopProductStatusResult)
}

// ReleaseTopProductStatusResult 释放TopProductStatusResult
func ReleaseTopProductStatusResult(v *TopProductStatusResult) {
	v.Release = false
	poolTopProductStatusResult.Put(v)
}
