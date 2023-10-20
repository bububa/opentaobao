package aliexpresssumaitong

import (
	"sync"
)

// Item 结构体
type Item struct {
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品价格
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
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
	v.ItemId = 0
	v.Price = 0
	v.Quantity = 0
	v.SkuId = 0
	poolItem.Put(v)
}
