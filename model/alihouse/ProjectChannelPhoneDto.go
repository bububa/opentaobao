package alihouse

import (
	"sync"
)

// ProjectChannelPhoneDto 结构体
type ProjectChannelPhoneDto struct {
	// 楼盘outerid
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 渠道电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 批次号
	BatchNo string `json:"batch_no,omitempty" xml:"batch_no,omitempty"`
	// 渠道类型 1-高德
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 是否删除 1-已删除 0-未删除
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}

var poolProjectChannelPhoneDto = sync.Pool{
	New: func() any {
		return new(ProjectChannelPhoneDto)
	},
}

// GetProjectChannelPhoneDto() 从对象池中获取ProjectChannelPhoneDto
func GetProjectChannelPhoneDto() *ProjectChannelPhoneDto {
	return poolProjectChannelPhoneDto.Get().(*ProjectChannelPhoneDto)
}

// ReleaseProjectChannelPhoneDto 释放ProjectChannelPhoneDto
func ReleaseProjectChannelPhoneDto(v *ProjectChannelPhoneDto) {
	v.OuterId = ""
	v.Phone = ""
	v.BatchNo = ""
	v.Channel = 0
	v.IsDeleted = 0
	poolProjectChannelPhoneDto.Put(v)
}
