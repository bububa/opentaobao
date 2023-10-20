package jstinteractive

import (
	"sync"
)

// Progress 结构体
type Progress struct {
	// 任务的完成最大上限次数
	MaxTimes int64 `json:"max_times,omitempty" xml:"max_times,omitempty"`
	// 用户已完成的任务次数
	Times int64 `json:"times,omitempty" xml:"times,omitempty"`
	// 用户还需要完成的任务次数，needTimes=maxTimes-times
	NeedTimes int64 `json:"need_times,omitempty" xml:"need_times,omitempty"`
}

var poolProgress = sync.Pool{
	New: func() any {
		return new(Progress)
	},
}

// GetProgress() 从对象池中获取Progress
func GetProgress() *Progress {
	return poolProgress.Get().(*Progress)
}

// ReleaseProgress 释放Progress
func ReleaseProgress(v *Progress) {
	v.MaxTimes = 0
	v.Times = 0
	v.NeedTimes = 0
	poolProgress.Put(v)
}
