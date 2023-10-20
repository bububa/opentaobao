package servicecenter

import (
	"sync"
)

// CsSchedulingWrapper 结构体
type CsSchedulingWrapper struct {
	// 按天排班信息
	CsSchedulings []CsScheduling `json:"cs_schedulings,omitempty" xml:"cs_schedulings>cs_scheduling,omitempty"`
	// 排班记录条数（按天计算）
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolCsSchedulingWrapper = sync.Pool{
	New: func() any {
		return new(CsSchedulingWrapper)
	},
}

// GetCsSchedulingWrapper() 从对象池中获取CsSchedulingWrapper
func GetCsSchedulingWrapper() *CsSchedulingWrapper {
	return poolCsSchedulingWrapper.Get().(*CsSchedulingWrapper)
}

// ReleaseCsSchedulingWrapper 释放CsSchedulingWrapper
func ReleaseCsSchedulingWrapper(v *CsSchedulingWrapper) {
	v.CsSchedulings = v.CsSchedulings[:0]
	v.TotalCount = 0
	poolCsSchedulingWrapper.Put(v)
}
