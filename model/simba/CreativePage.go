package simba

import (
	"sync"
)

// CreativePage 结构体
type CreativePage struct {
	// 广告创意列表
	CreativeList []Creative `json:"creative_list,omitempty" xml:"creative_list>creative,omitempty"`
	// 每页数据大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 返回第几页的数据从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}

var poolCreativePage = sync.Pool{
	New: func() any {
		return new(CreativePage)
	},
}

// GetCreativePage() 从对象池中获取CreativePage
func GetCreativePage() *CreativePage {
	return poolCreativePage.Get().(*CreativePage)
}

// ReleaseCreativePage 释放CreativePage
func ReleaseCreativePage(v *CreativePage) {
	v.CreativeList = v.CreativeList[:0]
	v.PageSize = 0
	v.PageNo = 0
	v.TotalItem = 0
	poolCreativePage.Put(v)
}
