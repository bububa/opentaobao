package opentrade

import (
	"sync"
)

// McSceneActivityDto 结构体
type McSceneActivityDto struct {
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动关联的商品列表，使用逗号(,)分割
	ItemIds string `json:"item_ids,omitempty" xml:"item_ids,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 创建活动的appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 排队活动ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 最近修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolMcSceneActivityDto = sync.Pool{
	New: func() any {
		return new(McSceneActivityDto)
	},
}

// GetMcSceneActivityDto() 从对象池中获取McSceneActivityDto
func GetMcSceneActivityDto() *McSceneActivityDto {
	return poolMcSceneActivityDto.Get().(*McSceneActivityDto)
}

// ReleaseMcSceneActivityDto 释放McSceneActivityDto
func ReleaseMcSceneActivityDto(v *McSceneActivityDto) {
	v.StartTime = ""
	v.ItemIds = ""
	v.ActivityName = ""
	v.Appkey = ""
	v.ActivityId = ""
	v.GmtModified = ""
	v.EndTime = ""
	poolMcSceneActivityDto.Put(v)
}
