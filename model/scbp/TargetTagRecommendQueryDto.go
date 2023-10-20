package scbp

import (
	"sync"
)

// TargetTagRecommendQueryDto 结构体
type TargetTagRecommendQueryDto struct {
	// 查询类型(查询推荐标签丰富信息，包含拓展信息和标签美杜莎文案（DEFAULT） 查询推荐标签扩展信息，包含基础信息和算法推荐溢价值（SIMPLE） 查询推荐标签基础信息（RAW）)
	QueryMode string `json:"query_mode,omitempty" xml:"query_mode,omitempty"`
	// 是否需要过滤已采纳标签
	IsFilterAdoptedTag bool `json:"is_filter_adopted_tag,omitempty" xml:"is_filter_adopted_tag,omitempty"`
}

var poolTargetTagRecommendQueryDto = sync.Pool{
	New: func() any {
		return new(TargetTagRecommendQueryDto)
	},
}

// GetTargetTagRecommendQueryDto() 从对象池中获取TargetTagRecommendQueryDto
func GetTargetTagRecommendQueryDto() *TargetTagRecommendQueryDto {
	return poolTargetTagRecommendQueryDto.Get().(*TargetTagRecommendQueryDto)
}

// ReleaseTargetTagRecommendQueryDto 释放TargetTagRecommendQueryDto
func ReleaseTargetTagRecommendQueryDto(v *TargetTagRecommendQueryDto) {
	v.QueryMode = ""
	v.IsFilterAdoptedTag = false
	poolTargetTagRecommendQueryDto.Put(v)
}
