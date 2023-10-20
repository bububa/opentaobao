package idleisv

import (
	"sync"
)

// IsvRefundTimeoutDto 结构体
type IsvRefundTimeoutDto struct {
	// 退款超时动作类型
	TimeoutActionType string `json:"timeout_action_type,omitempty" xml:"timeout_action_type,omitempty"`
	// 退款超时创建时间，时间戳，毫秒
	Create int64 `json:"create,omitempty" xml:"create,omitempty"`
	// 退款超时修改时间，时间戳，毫秒
	Modified int64 `json:"modified,omitempty" xml:"modified,omitempty"`
	// 退款超时时长，时间长度(毫秒)
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 退款超时时间点(自动触发状态变更的超时时间点)，时间戳，毫秒
	Timeout int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 退款超时运行状态 0:超时创建完成, 1:超时运行中, 2:超时暂停, 3:超时关闭, 4:超时失败, 5:超时成功
	TimeoutStatus int64 `json:"timeout_status,omitempty" xml:"timeout_status,omitempty"`
}

var poolIsvRefundTimeoutDto = sync.Pool{
	New: func() any {
		return new(IsvRefundTimeoutDto)
	},
}

// GetIsvRefundTimeoutDto() 从对象池中获取IsvRefundTimeoutDto
func GetIsvRefundTimeoutDto() *IsvRefundTimeoutDto {
	return poolIsvRefundTimeoutDto.Get().(*IsvRefundTimeoutDto)
}

// ReleaseIsvRefundTimeoutDto 释放IsvRefundTimeoutDto
func ReleaseIsvRefundTimeoutDto(v *IsvRefundTimeoutDto) {
	v.TimeoutActionType = ""
	v.Create = 0
	v.Modified = 0
	v.Duration = 0
	v.Timeout = 0
	v.TimeoutStatus = 0
	poolIsvRefundTimeoutDto.Put(v)
}
