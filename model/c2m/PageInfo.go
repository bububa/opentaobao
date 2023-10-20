package c2m

import (
	"sync"
)

// PageInfo 结构体
type PageInfo struct {
	// 邀请关系信息
	List []OrganizationInviteInfoVo `json:"list,omitempty" xml:"list>organization_invite_info_vo,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 总页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 第几页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
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
	v.Total = 0
	v.Pages = 0
	v.PageSize = 0
	v.PageNum = 0
	poolPageInfo.Put(v)
}
