package promotion

import (
	"sync"
)

// PromotionTagQuery 结构体
type PromotionTagQuery struct {
	// 标签结果列表
	TagList []PromotionTag `json:"tag_list,omitempty" xml:"tag_list>promotion_tag,omitempty"`
	// 总记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

var poolPromotionTagQuery = sync.Pool{
	New: func() any {
		return new(PromotionTagQuery)
	},
}

// GetPromotionTagQuery() 从对象池中获取PromotionTagQuery
func GetPromotionTagQuery() *PromotionTagQuery {
	return poolPromotionTagQuery.Get().(*PromotionTagQuery)
}

// ReleasePromotionTagQuery 释放PromotionTagQuery
func ReleasePromotionTagQuery(v *PromotionTagQuery) {
	v.TagList = v.TagList[:0]
	v.TotalResults = 0
	poolPromotionTagQuery.Put(v)
}
