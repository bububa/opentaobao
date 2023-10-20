package promotion

import (
	"sync"
)

// PageInfo 结构体
type PageInfo struct {
	// 第几页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 每页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总共多少页
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 总共多少条
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	v.PageNum = 0
	v.PageSize = 0
	v.Pages = 0
	v.Total = 0
	poolPageInfo.Put(v)
}
