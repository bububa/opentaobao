package tblogistics

import (
	"sync"
)

// TransitStepInfo 结构体
type TransitStepInfo struct {
	// 状态发生的时间
	StatusTime string `json:"status_time,omitempty" xml:"status_time,omitempty"`
	// 状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 节点说明 ，指明当前节点揽收、派送，签收。
	Action string `json:"action,omitempty" xml:"action,omitempty"`
}

var poolTransitStepInfo = sync.Pool{
	New: func() any {
		return new(TransitStepInfo)
	},
}

// GetTransitStepInfo() 从对象池中获取TransitStepInfo
func GetTransitStepInfo() *TransitStepInfo {
	return poolTransitStepInfo.Get().(*TransitStepInfo)
}

// ReleaseTransitStepInfo 释放TransitStepInfo
func ReleaseTransitStepInfo(v *TransitStepInfo) {
	v.StatusTime = ""
	v.StatusDesc = ""
	v.Action = ""
	poolTransitStepInfo.Put(v)
}
