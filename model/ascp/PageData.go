package ascp

import (
	"sync"
)

// PageData 结构体
type PageData struct {
	// 发货单列表
	List []ConsignOrder `json:"list,omitempty" xml:"list>consign_order,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页数量，不大于100
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolPageData = sync.Pool{
	New: func() any {
		return new(PageData)
	},
}

// GetPageData() 从对象池中获取PageData
func GetPageData() *PageData {
	return poolPageData.Get().(*PageData)
}

// ReleasePageData 释放PageData
func ReleasePageData(v *PageData) {
	v.List = v.List[:0]
	v.PageNo = 0
	v.PageSize = 0
	v.HasNext = false
	poolPageData.Put(v)
}
