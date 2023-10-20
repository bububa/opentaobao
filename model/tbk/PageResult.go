package tbk

import (
	"sync"
)

// PageResult 结构体
type PageResult struct {
	// 数据结果
	Results []TaobaoTbkDgCpaActivityDetailResults `json:"results,omitempty" xml:"results>taobao_tbk_dg_cpa_activity_detail_results,omitempty"`
	// 数据批次号(时间戳)
	Runtime string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 上一页页码
	PrePage int64 `json:"pre_page,omitempty" xml:"pre_page,omitempty"`
	// 下一页页码
	NextPage int64 `json:"next_page,omitempty" xml:"next_page,omitempty"`
	// 当前页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 总共页数
	TotalPages int64 `json:"total_pages,omitempty" xml:"total_pages,omitempty"`
	// 每页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 是否有下一页
	HasPre bool `json:"has_pre,omitempty" xml:"has_pre,omitempty"`
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
	v.Runtime = ""
	v.PrePage = 0
	v.NextPage = 0
	v.PageNo = 0
	v.TotalPages = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.HasNext = false
	v.HasPre = false
	poolPageResult.Put(v)
}
