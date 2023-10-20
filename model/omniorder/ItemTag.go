package omniorder

import (
	"sync"
)

// ItemTag 结构体
type ItemTag struct {
	// tagType
	TagType string `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolItemTag = sync.Pool{
	New: func() any {
		return new(ItemTag)
	},
}

// GetItemTag() 从对象池中获取ItemTag
func GetItemTag() *ItemTag {
	return poolItemTag.Get().(*ItemTag)
}

// ReleaseItemTag 释放ItemTag
func ReleaseItemTag(v *ItemTag) {
	v.TagType = ""
	v.ItemId = 0
	poolItemTag.Put(v)
}
