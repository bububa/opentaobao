package alihealth2

import (
	"sync"
)

// TopChannelDeptSyncDto 结构体
type TopChannelDeptSyncDto struct {
	// 医院ID+科室ID+状态
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
}

var poolTopChannelDeptSyncDto = sync.Pool{
	New: func() any {
		return new(TopChannelDeptSyncDto)
	},
}

// GetTopChannelDeptSyncDto() 从对象池中获取TopChannelDeptSyncDto
func GetTopChannelDeptSyncDto() *TopChannelDeptSyncDto {
	return poolTopChannelDeptSyncDto.Get().(*TopChannelDeptSyncDto)
}

// ReleaseTopChannelDeptSyncDto 释放TopChannelDeptSyncDto
func ReleaseTopChannelDeptSyncDto(v *TopChannelDeptSyncDto) {
	v.BizContent = ""
	poolTopChannelDeptSyncDto.Put(v)
}
