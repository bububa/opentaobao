package examination

import (
	"sync"
)

// Item 结构体
type Item struct {
	// 体检标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 体检项描述
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 体检项编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 体检组编码
	ItemGroupCode string `json:"item_group_code,omitempty" xml:"item_group_code,omitempty"`
	// 体检组显示权重
	ItemGroupWeight string `json:"item_group_weight,omitempty" xml:"item_group_weight,omitempty"`
}

var poolItem = sync.Pool{
	New: func() any {
		return new(Item)
	},
}

// GetItem() 从对象池中获取Item
func GetItem() *Item {
	return poolItem.Get().(*Item)
}

// ReleaseItem 释放Item
func ReleaseItem(v *Item) {
	v.Title = ""
	v.Detail = ""
	v.SubTitle = ""
	v.ItemCode = ""
	v.ItemGroupCode = ""
	v.ItemGroupWeight = ""
	poolItem.Put(v)
}
