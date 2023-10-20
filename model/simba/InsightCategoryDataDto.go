package simba

import (
	"sync"
)

// InsightCategoryDataDto 结构体
type InsightCategoryDataDto struct {
	// 类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 展现量
	Impression int64 `json:"impression,omitempty" xml:"impression,omitempty"`
	// 点击量
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}

var poolInsightCategoryDataDto = sync.Pool{
	New: func() any {
		return new(InsightCategoryDataDto)
	},
}

// GetInsightCategoryDataDto() 从对象池中获取InsightCategoryDataDto
func GetInsightCategoryDataDto() *InsightCategoryDataDto {
	return poolInsightCategoryDataDto.Get().(*InsightCategoryDataDto)
}

// ReleaseInsightCategoryDataDto 释放InsightCategoryDataDto
func ReleaseInsightCategoryDataDto(v *InsightCategoryDataDto) {
	v.CatName = ""
	v.Ctr = ""
	v.Impression = 0
	v.Click = 0
	v.CatId = 0
	poolInsightCategoryDataDto.Put(v)
}
