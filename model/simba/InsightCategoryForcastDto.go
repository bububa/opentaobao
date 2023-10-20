package simba

import (
	"sync"
)

// InsightCategoryForcastDto 结构体
type InsightCategoryForcastDto struct {
	// 表示类目的所有父级类目的名称，不同层级之间用（ascii码为2的字符）分隔开
	CatPathName string `json:"cat_path_name,omitempty" xml:"cat_path_name,omitempty"`
	// 表示词与该类目的相关度，值越大表示越相关
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 表示该类目的所有父级类目，按层级顺序排列，层级越高的在前面，不同的层级之间用空格分隔
	CatPathId string `json:"cat_path_id,omitempty" xml:"cat_path_id,omitempty"`
	// 查询词
	Bidword string `json:"bidword,omitempty" xml:"bidword,omitempty"`
}

var poolInsightCategoryForcastDto = sync.Pool{
	New: func() any {
		return new(InsightCategoryForcastDto)
	},
}

// GetInsightCategoryForcastDto() 从对象池中获取InsightCategoryForcastDto
func GetInsightCategoryForcastDto() *InsightCategoryForcastDto {
	return poolInsightCategoryForcastDto.Get().(*InsightCategoryForcastDto)
}

// ReleaseInsightCategoryForcastDto 释放InsightCategoryForcastDto
func ReleaseInsightCategoryForcastDto(v *InsightCategoryForcastDto) {
	v.CatPathName = ""
	v.Score = ""
	v.CatPathId = ""
	v.Bidword = ""
	poolInsightCategoryForcastDto.Put(v)
}
