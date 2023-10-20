package alihealth2

import (
	"sync"
)

// PageInfo 结构体
type PageInfo struct {
	// result
	Results []AlibabaAlihealthTracecodesellerChannelSearchResult `json:"results,omitempty" xml:"results>alibaba_alihealth_tracecodeseller_channel_search_result,omitempty"`
	// totalNum
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// page
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPageInfo = sync.Pool{
	New: func() any {
		return new(PageInfo)
	},
}

// GetPageInfo() 从对象池中获取PageInfo
func GetPageInfo() *PageInfo {
	return poolPageInfo.Get().(*PageInfo)
}

// ReleasePageInfo 释放PageInfo
func ReleasePageInfo(v *PageInfo) {
	v.Results = v.Results[:0]
	v.TotalNum = 0
	v.Page = 0
	v.PageSize = 0
	poolPageInfo.Put(v)
}
