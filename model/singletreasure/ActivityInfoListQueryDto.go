package singletreasure

import (
	"sync"
)

// ActivityInfoListQueryDto 结构体
type ActivityInfoListQueryDto struct {
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动 id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动状态  删除：-1；暂停：0；未开始：1；进行中：2；已结束：3
	ActivityStatus int64 `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
}

var poolActivityInfoListQueryDto = sync.Pool{
	New: func() any {
		return new(ActivityInfoListQueryDto)
	},
}

// GetActivityInfoListQueryDto() 从对象池中获取ActivityInfoListQueryDto
func GetActivityInfoListQueryDto() *ActivityInfoListQueryDto {
	return poolActivityInfoListQueryDto.Get().(*ActivityInfoListQueryDto)
}

// ReleaseActivityInfoListQueryDto 释放ActivityInfoListQueryDto
func ReleaseActivityInfoListQueryDto(v *ActivityInfoListQueryDto) {
	v.Name = ""
	v.ActivityId = 0
	v.ActivityStatus = 0
	poolActivityInfoListQueryDto.Put(v)
}
