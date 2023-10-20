package charity

import (
	"sync"
)

// SumUserCharityTimeTopApiHsfResponse 结构体
type SumUserCharityTimeTopApiHsfResponse struct {
	// 公益时：分钟
	CharityTimeMinutes int64 `json:"charity_time_minutes,omitempty" xml:"charity_time_minutes,omitempty"`
	// 公益时：小时
	CharityTimeHours float64 `json:"charity_time_hours,omitempty" xml:"charity_time_hours,omitempty"`
}

var poolSumUserCharityTimeTopApiHsfResponse = sync.Pool{
	New: func() any {
		return new(SumUserCharityTimeTopApiHsfResponse)
	},
}

// GetSumUserCharityTimeTopApiHsfResponse() 从对象池中获取SumUserCharityTimeTopApiHsfResponse
func GetSumUserCharityTimeTopApiHsfResponse() *SumUserCharityTimeTopApiHsfResponse {
	return poolSumUserCharityTimeTopApiHsfResponse.Get().(*SumUserCharityTimeTopApiHsfResponse)
}

// ReleaseSumUserCharityTimeTopApiHsfResponse 释放SumUserCharityTimeTopApiHsfResponse
func ReleaseSumUserCharityTimeTopApiHsfResponse(v *SumUserCharityTimeTopApiHsfResponse) {
	v.CharityTimeMinutes = 0
	v.CharityTimeHours = 0
	poolSumUserCharityTimeTopApiHsfResponse.Put(v)
}
