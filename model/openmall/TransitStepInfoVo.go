package openmall

import (
	"sync"
)

// TransitStepInfoVo 结构体
type TransitStepInfoVo struct {
	// 节点说明
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 状态发生的时间
	StatusTime string `json:"status_time,omitempty" xml:"status_time,omitempty"`
}

var poolTransitStepInfoVo = sync.Pool{
	New: func() any {
		return new(TransitStepInfoVo)
	},
}

// GetTransitStepInfoVo() 从对象池中获取TransitStepInfoVo
func GetTransitStepInfoVo() *TransitStepInfoVo {
	return poolTransitStepInfoVo.Get().(*TransitStepInfoVo)
}

// ReleaseTransitStepInfoVo 释放TransitStepInfoVo
func ReleaseTransitStepInfoVo(v *TransitStepInfoVo) {
	v.Action = ""
	v.StatusDesc = ""
	v.StatusTime = ""
	poolTransitStepInfoVo.Put(v)
}
