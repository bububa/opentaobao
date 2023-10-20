package ascpchannel

import (
	"sync"
)

// Futureplaninfodto 结构体
type Futureplaninfodto struct {
	// 销售结束时间，YYYYMMDDHHMMSS
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 销售开始时间，YYYYMMDDHHMMSS
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 默认填3
	AicFutureInvPublishType int64 `json:"aic_future_inv_publish_type,omitempty" xml:"aic_future_inv_publish_type,omitempty"`
	// 默认填1
	AicFutureInvTimeStrategy int64 `json:"aic_future_inv_time_strategy,omitempty" xml:"aic_future_inv_time_strategy,omitempty"`
	// 相对时间天数(单位:天 适用于相对计划)
	RelativeTime int64 `json:"relative_time,omitempty" xml:"relative_time,omitempty"`
}

var poolFutureplaninfodto = sync.Pool{
	New: func() any {
		return new(Futureplaninfodto)
	},
}

// GetFutureplaninfodto() 从对象池中获取Futureplaninfodto
func GetFutureplaninfodto() *Futureplaninfodto {
	return poolFutureplaninfodto.Get().(*Futureplaninfodto)
}

// ReleaseFutureplaninfodto 释放Futureplaninfodto
func ReleaseFutureplaninfodto(v *Futureplaninfodto) {
	v.EndDate = ""
	v.StartDate = ""
	v.AicFutureInvPublishType = 0
	v.AicFutureInvTimeStrategy = 0
	v.RelativeTime = 0
	poolFutureplaninfodto.Put(v)
}
