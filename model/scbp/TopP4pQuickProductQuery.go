package scbp

import (
	"sync"
)

// TopP4pQuickProductQuery 结构体
type TopP4pQuickProductQuery struct {
	// 定向推广计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
	// 每页返回多少行
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
}

var poolTopP4pQuickProductQuery = sync.Pool{
	New: func() any {
		return new(TopP4pQuickProductQuery)
	},
}

// GetTopP4pQuickProductQuery() 从对象池中获取TopP4pQuickProductQuery
func GetTopP4pQuickProductQuery() *TopP4pQuickProductQuery {
	return poolTopP4pQuickProductQuery.Get().(*TopP4pQuickProductQuery)
}

// ReleaseTopP4pQuickProductQuery 释放TopP4pQuickProductQuery
func ReleaseTopP4pQuickProductQuery(v *TopP4pQuickProductQuery) {
	v.CampaignId = 0
	v.ToPage = 0
	v.PerPageSize = 0
	poolTopP4pQuickProductQuery.Put(v)
}
