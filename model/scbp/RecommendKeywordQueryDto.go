package scbp

import (
	"sync"
)

// RecommendKeywordQueryDto 结构体
type RecommendKeywordQueryDto struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 排序字段：searchIndex,buyIndex,star,keyword
	OrderBy string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// 排序规则 asc/desc
	Order string `json:"order,omitempty" xml:"order,omitempty"`
	// 搜索类型：1-系统推荐                      2-按词分组推荐                      3-按词推荐词
	SearchType int64 `json:"search_type,omitempty" xml:"search_type,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 词组id，searchType为2时必传
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// true-精准匹配 false-模糊匹配
	ExactMatch bool `json:"exact_match,omitempty" xml:"exact_match,omitempty"`
}

var poolRecommendKeywordQueryDto = sync.Pool{
	New: func() any {
		return new(RecommendKeywordQueryDto)
	},
}

// GetRecommendKeywordQueryDto() 从对象池中获取RecommendKeywordQueryDto
func GetRecommendKeywordQueryDto() *RecommendKeywordQueryDto {
	return poolRecommendKeywordQueryDto.Get().(*RecommendKeywordQueryDto)
}

// ReleaseRecommendKeywordQueryDto 释放RecommendKeywordQueryDto
func ReleaseRecommendKeywordQueryDto(v *RecommendKeywordQueryDto) {
	v.Keyword = ""
	v.OrderBy = ""
	v.Order = ""
	v.SearchType = 0
	v.CampaignId = 0
	v.GroupId = 0
	v.ExactMatch = false
	poolRecommendKeywordQueryDto.Put(v)
}
