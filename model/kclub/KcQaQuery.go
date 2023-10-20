package kclub

import (
	"sync"
)

// KcQaQuery 结构体
type KcQaQuery struct {
	// context列表
	ContextList []int64 `json:"context_list,omitempty" xml:"context_list>int64,omitempty"`
	// 问题类型列表
	QuestionTypes []string `json:"question_types,omitempty" xml:"question_types>string,omitempty"`
	// 状态列表
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 视角过滤
	Views []string `json:"views,omitempty" xml:"views>string,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// context
	Context int64 `json:"context,omitempty" xml:"context,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 问题类型
	QuestionType int64 `json:"question_type,omitempty" xml:"question_type,omitempty"`
	// 排序对象
	SorterConfig *SorterConfig `json:"sorter_config,omitempty" xml:"sorter_config,omitempty"`
}

var poolKcQaQuery = sync.Pool{
	New: func() any {
		return new(KcQaQuery)
	},
}

// GetKcQaQuery() 从对象池中获取KcQaQuery
func GetKcQaQuery() *KcQaQuery {
	return poolKcQaQuery.Get().(*KcQaQuery)
}

// ReleaseKcQaQuery 释放KcQaQuery
func ReleaseKcQaQuery(v *KcQaQuery) {
	v.ContextList = v.ContextList[:0]
	v.QuestionTypes = v.QuestionTypes[:0]
	v.StatusList = v.StatusList[:0]
	v.Views = v.Views[:0]
	v.TenantId = 0
	v.Status = 0
	v.Context = 0
	v.PageSize = 0
	v.CatId = 0
	v.CurrentPage = 0
	v.QuestionType = 0
	v.SorterConfig = nil
	poolKcQaQuery.Put(v)
}
