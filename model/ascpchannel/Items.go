package ascpchannel

import (
	"sync"
)

// Items 结构体
type Items struct {
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商品仓储系统编码
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 包裹内该商品的数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolItems = sync.Pool{
	New: func() any {
		return new(Items)
	},
}

// GetItems() 从对象池中获取Items
func GetItems() *Items {
	return poolItems.Get().(*Items)
}

// ReleaseItems 释放Items
func ReleaseItems(v *Items) {
	v.ItemCode = ""
	v.ItemId = 0
	v.Quantity = 0
	poolItems.Put(v)
}
