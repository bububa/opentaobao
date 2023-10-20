package alsc

import (
	"sync"
)

// CrmTagOpenInfo 结构体
type CrmTagOpenInfo struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 会员计划ID
	PlanId string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// 标签ID
	TagId string `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 标签名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 标签类型
	TagType string `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
	// 是否已删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolCrmTagOpenInfo = sync.Pool{
	New: func() any {
		return new(CrmTagOpenInfo)
	},
}

// GetCrmTagOpenInfo() 从对象池中获取CrmTagOpenInfo
func GetCrmTagOpenInfo() *CrmTagOpenInfo {
	return poolCrmTagOpenInfo.Get().(*CrmTagOpenInfo)
}

// ReleaseCrmTagOpenInfo 释放CrmTagOpenInfo
func ReleaseCrmTagOpenInfo(v *CrmTagOpenInfo) {
	v.GmtCreate = ""
	v.GmtModified = ""
	v.PlanId = ""
	v.TagId = ""
	v.TagName = ""
	v.TagType = ""
	v.Deleted = false
	poolCrmTagOpenInfo.Put(v)
}
