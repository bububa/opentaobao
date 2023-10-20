package tvupadmin

import (
	"sync"
)

// PageResult 结构体
type PageResult struct {
	// 技能列表
	Results []SkillSimpleView `json:"results,omitempty" xml:"results>skill_simple_view,omitempty"`
	// 总页数
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 分页单位
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolPageResult = sync.Pool{
	New: func() any {
		return new(PageResult)
	},
}

// GetPageResult() 从对象池中获取PageResult
func GetPageResult() *PageResult {
	return poolPageResult.Get().(*PageResult)
}

// ReleasePageResult 释放PageResult
func ReleasePageResult(v *PageResult) {
	v.Results = v.Results[:0]
	v.PageCount = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.CurrentPage = 0
	poolPageResult.Put(v)
}
