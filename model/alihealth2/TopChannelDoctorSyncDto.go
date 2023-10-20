package alihealth2

import (
	"sync"
)

// TopChannelDoctorSyncDto 结构体
type TopChannelDoctorSyncDto struct {
	// 医生下架参数
	BizContent string `json:"biz_content,omitempty" xml:"biz_content,omitempty"`
}

var poolTopChannelDoctorSyncDto = sync.Pool{
	New: func() any {
		return new(TopChannelDoctorSyncDto)
	},
}

// GetTopChannelDoctorSyncDto() 从对象池中获取TopChannelDoctorSyncDto
func GetTopChannelDoctorSyncDto() *TopChannelDoctorSyncDto {
	return poolTopChannelDoctorSyncDto.Get().(*TopChannelDoctorSyncDto)
}

// ReleaseTopChannelDoctorSyncDto 释放TopChannelDoctorSyncDto
func ReleaseTopChannelDoctorSyncDto(v *TopChannelDoctorSyncDto) {
	v.BizContent = ""
	poolTopChannelDoctorSyncDto.Put(v)
}
