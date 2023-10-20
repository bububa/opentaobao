package examination

import (
	"sync"
)

// AddItem 结构体
type AddItem struct {
	// 加项id
	IsvItemId string `json:"isv_item_id,omitempty" xml:"isv_item_id,omitempty"`
	// 版本号，isv需要进行校验版本号是否过期，判断加项信息是否已更改，健康未同步
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolAddItem = sync.Pool{
	New: func() any {
		return new(AddItem)
	},
}

// GetAddItem() 从对象池中获取AddItem
func GetAddItem() *AddItem {
	return poolAddItem.Get().(*AddItem)
}

// ReleaseAddItem 释放AddItem
func ReleaseAddItem(v *AddItem) {
	v.IsvItemId = ""
	v.Version = 0
	poolAddItem.Put(v)
}
