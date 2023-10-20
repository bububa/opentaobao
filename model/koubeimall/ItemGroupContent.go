package koubeimall

import (
	"sync"
)

// ItemGroupContent 结构体
type ItemGroupContent struct {
	// 详情组列表
	ContentGroupDetailList []ItemGroupContentDetail `json:"content_group_detail_list,omitempty" xml:"content_group_detail_list>item_group_content_detail,omitempty"`
	// 组名
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

var poolItemGroupContent = sync.Pool{
	New: func() any {
		return new(ItemGroupContent)
	},
}

// GetItemGroupContent() 从对象池中获取ItemGroupContent
func GetItemGroupContent() *ItemGroupContent {
	return poolItemGroupContent.Get().(*ItemGroupContent)
}

// ReleaseItemGroupContent 释放ItemGroupContent
func ReleaseItemGroupContent(v *ItemGroupContent) {
	v.ContentGroupDetailList = v.ContentGroupDetailList[:0]
	v.GroupName = ""
	poolItemGroupContent.Put(v)
}
