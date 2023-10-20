package iot

import (
	"sync"
)

// PageInfo 结构体
type PageInfo struct {
	// 数据集
	List []BusinessRecipeOpenDto `json:"list,omitempty" xml:"list>business_recipe_open_dto,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 总条数
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
	v.List = v.List[:0]
	v.PageSize = 0
	v.PageNum = 0
	v.Total = 0
	poolPageInfo.Put(v)
}
