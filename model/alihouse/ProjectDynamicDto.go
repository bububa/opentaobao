package alihouse

import (
	"sync"
)

// ProjectDynamicDto 结构体
type ProjectDynamicDto struct {
	// 楼盘标签
	ProjectTags []string `json:"project_tags,omitempty" xml:"project_tags>string,omitempty"`
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部快讯ID
	OuterDynamicId string `json:"outer_dynamic_id,omitempty" xml:"outer_dynamic_id,omitempty"`
	// 快讯标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 快讯正文
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 发布时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 状态 0-无效 1-有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolProjectDynamicDto = sync.Pool{
	New: func() any {
		return new(ProjectDynamicDto)
	},
}

// GetProjectDynamicDto() 从对象池中获取ProjectDynamicDto
func GetProjectDynamicDto() *ProjectDynamicDto {
	return poolProjectDynamicDto.Get().(*ProjectDynamicDto)
}

// ReleaseProjectDynamicDto 释放ProjectDynamicDto
func ReleaseProjectDynamicDto(v *ProjectDynamicDto) {
	v.ProjectTags = v.ProjectTags[:0]
	v.OuterId = ""
	v.OuterDynamicId = ""
	v.Title = ""
	v.Content = ""
	v.PublishTime = ""
	v.OuterStoreId = ""
	v.CityId = 0
	v.Status = 0
	poolProjectDynamicDto.Put(v)
}
