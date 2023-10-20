package ascpchannel

import (
	"sync"
)

// Itemdolist 结构体
type Itemdolist struct {
	// 前端商品 ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 前端SKU ID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolItemdolist = sync.Pool{
	New: func() any {
		return new(Itemdolist)
	},
}

// GetItemdolist() 从对象池中获取Itemdolist
func GetItemdolist() *Itemdolist {
	return poolItemdolist.Get().(*Itemdolist)
}

// ReleaseItemdolist 释放Itemdolist
func ReleaseItemdolist(v *Itemdolist) {
	v.ItemId = ""
	v.SkuId = ""
	poolItemdolist.Put(v)
}
