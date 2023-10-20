package product

import (
	"sync"
)

// DapeiTemplateItem 结构体
type DapeiTemplateItem struct {
	// img
	Img string `json:"img,omitempty" xml:"img,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolDapeiTemplateItem = sync.Pool{
	New: func() any {
		return new(DapeiTemplateItem)
	},
}

// GetDapeiTemplateItem() 从对象池中获取DapeiTemplateItem
func GetDapeiTemplateItem() *DapeiTemplateItem {
	return poolDapeiTemplateItem.Get().(*DapeiTemplateItem)
}

// ReleaseDapeiTemplateItem 释放DapeiTemplateItem
func ReleaseDapeiTemplateItem(v *DapeiTemplateItem) {
	v.Img = ""
	v.ItemId = 0
	poolDapeiTemplateItem.Put(v)
}
