package logistic

import (
	"sync"
)

// ReachableDto 结构体
type ReachableDto struct {
	// 阻断原因
	InterruptReason string `json:"interrupt_reason,omitempty" xml:"interrupt_reason,omitempty"`
	// 是否阻断 true:阻断  false:可达
	InterruptApplyWaybillCode bool `json:"interrupt_apply_waybill_code,omitempty" xml:"interrupt_apply_waybill_code,omitempty"`
}

var poolReachableDto = sync.Pool{
	New: func() any {
		return new(ReachableDto)
	},
}

// GetReachableDto() 从对象池中获取ReachableDto
func GetReachableDto() *ReachableDto {
	return poolReachableDto.Get().(*ReachableDto)
}

// ReleaseReachableDto 释放ReachableDto
func ReleaseReachableDto(v *ReachableDto) {
	v.InterruptReason = ""
	v.InterruptApplyWaybillCode = false
	poolReachableDto.Put(v)
}
