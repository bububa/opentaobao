package aetask

import (
	"sync"
)

// QueryDeliveryRequestDto 结构体
type QueryDeliveryRequestDto struct {
	// 用户版本信息
	Ttid string `json:"ttid,omitempty" xml:"ttid,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 准入key
	ProjectAppKey string `json:"project_app_key,omitempty" xml:"project_app_key,omitempty"`
	// 0:不展示预热 1：展示预热
	PreDisplay int64 `json:"pre_display,omitempty" xml:"pre_display,omitempty"`
	// 投放场景id
	SceneId int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
}

var poolQueryDeliveryRequestDto = sync.Pool{
	New: func() any {
		return new(QueryDeliveryRequestDto)
	},
}

// GetQueryDeliveryRequestDto() 从对象池中获取QueryDeliveryRequestDto
func GetQueryDeliveryRequestDto() *QueryDeliveryRequestDto {
	return poolQueryDeliveryRequestDto.Get().(*QueryDeliveryRequestDto)
}

// ReleaseQueryDeliveryRequestDto 释放QueryDeliveryRequestDto
func ReleaseQueryDeliveryRequestDto(v *QueryDeliveryRequestDto) {
	v.Ttid = ""
	v.Language = ""
	v.Country = ""
	v.ProjectAppKey = ""
	v.PreDisplay = 0
	v.SceneId = 0
	poolQueryDeliveryRequestDto.Put(v)
}
