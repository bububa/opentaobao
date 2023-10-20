package maitix

import (
	"sync"
)

// StandQueryParam 结构体
type StandQueryParam struct {
	// 座位ID
	SeatIds []int64 `json:"seat_ids,omitempty" xml:"seat_ids>int64,omitempty"`
	// 看台ID
	StandId int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
}

var poolStandQueryParam = sync.Pool{
	New: func() any {
		return new(StandQueryParam)
	},
}

// GetStandQueryParam() 从对象池中获取StandQueryParam
func GetStandQueryParam() *StandQueryParam {
	return poolStandQueryParam.Get().(*StandQueryParam)
}

// ReleaseStandQueryParam 释放StandQueryParam
func ReleaseStandQueryParam(v *StandQueryParam) {
	v.SeatIds = v.SeatIds[:0]
	v.StandId = 0
	poolStandQueryParam.Put(v)
}
