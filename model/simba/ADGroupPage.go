package simba

import (
	"sync"
)

// ADGroupPage 结构体
type ADGroupPage struct {
	// 商品对象列表
	AdgroupList []ADGroup `json:"adgroup_list,omitempty" xml:"adgroup_list>ad_group,omitempty"`
	// 模板规则
	Schedule string `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// 模板名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 每页数据大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 返回第几页的数据从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 所查询的数据总数，只有当返回第一页数据时有值，当要求返回的数据非第一页时，此返回值无效
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
	// 模板id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolADGroupPage = sync.Pool{
	New: func() any {
		return new(ADGroupPage)
	},
}

// GetADGroupPage() 从对象池中获取ADGroupPage
func GetADGroupPage() *ADGroupPage {
	return poolADGroupPage.Get().(*ADGroupPage)
}

// ReleaseADGroupPage 释放ADGroupPage
func ReleaseADGroupPage(v *ADGroupPage) {
	v.AdgroupList = v.AdgroupList[:0]
	v.Schedule = ""
	v.Name = ""
	v.PageSize = 0
	v.PageNo = 0
	v.TotalItem = 0
	v.Id = 0
	poolADGroupPage.Put(v)
}
