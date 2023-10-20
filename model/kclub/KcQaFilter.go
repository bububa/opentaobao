package kclub

import (
	"sync"
)

// KcQaFilter 结构体
type KcQaFilter struct {
	// 视角
	Views []string `json:"views,omitempty" xml:"views>string,omitempty"`
	// 状态
	Statuses []string `json:"statuses,omitempty" xml:"statuses>string,omitempty"`
	// 是否需要子知识
	NeedChildKnowledge bool `json:"need_child_knowledge,omitempty" xml:"need_child_knowledge,omitempty"`
}

var poolKcQaFilter = sync.Pool{
	New: func() any {
		return new(KcQaFilter)
	},
}

// GetKcQaFilter() 从对象池中获取KcQaFilter
func GetKcQaFilter() *KcQaFilter {
	return poolKcQaFilter.Get().(*KcQaFilter)
}

// ReleaseKcQaFilter 释放KcQaFilter
func ReleaseKcQaFilter(v *KcQaFilter) {
	v.Views = v.Views[:0]
	v.Statuses = v.Statuses[:0]
	v.NeedChildKnowledge = false
	poolKcQaFilter.Put(v)
}
