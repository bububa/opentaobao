package scbp

import (
	"sync"
)

// ProductRecommendQueryDto 结构体
type ProductRecommendQueryDto struct {
	// TESTING：新发产品 HOT：优质转化品 WIN： 橱窗产品 PRE： 优先推广品 HIGH_VIDEO_PROD：优质视频品 PRODTAG_HOTSALE,PRODTAG_STRENGTH: 优品 SITE_NEW: 平台新品 DIRECT_PROD: 行业定征新品
	TagList []string `json:"tag_list,omitempty" xml:"tag_list>string,omitempty"`
	// 优品必填&#34;or&#34;
	TagQueryType string `json:"tag_query_type,omitempty" xml:"tag_query_type,omitempty"`
	// 新品必填：“NEW_PROD”
	SubType string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 当前页数
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 品创建天数,新品必传“90/180”
	CreateDays int64 `json:"create_days,omitempty" xml:"create_days,omitempty"`
}

var poolProductRecommendQueryDto = sync.Pool{
	New: func() any {
		return new(ProductRecommendQueryDto)
	},
}

// GetProductRecommendQueryDto() 从对象池中获取ProductRecommendQueryDto
func GetProductRecommendQueryDto() *ProductRecommendQueryDto {
	return poolProductRecommendQueryDto.Get().(*ProductRecommendQueryDto)
}

// ReleaseProductRecommendQueryDto 释放ProductRecommendQueryDto
func ReleaseProductRecommendQueryDto(v *ProductRecommendQueryDto) {
	v.TagList = v.TagList[:0]
	v.TagQueryType = ""
	v.SubType = ""
	v.Page = 0
	v.Size = 0
	v.CreateDays = 0
	poolProductRecommendQueryDto.Put(v)
}
