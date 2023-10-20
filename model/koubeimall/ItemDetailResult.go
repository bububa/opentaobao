package koubeimall

import (
	"sync"
)

// ItemDetailResult 结构体
type ItemDetailResult struct {
	// 商品详情信息
	ItemDetail *ItemDetailDto `json:"item_detail,omitempty" xml:"item_detail,omitempty"`
	// 商品基础信息
	ItemInfo *ItemDto `json:"item_info,omitempty" xml:"item_info,omitempty"`
}

var poolItemDetailResult = sync.Pool{
	New: func() any {
		return new(ItemDetailResult)
	},
}

// GetItemDetailResult() 从对象池中获取ItemDetailResult
func GetItemDetailResult() *ItemDetailResult {
	return poolItemDetailResult.Get().(*ItemDetailResult)
}

// ReleaseItemDetailResult 释放ItemDetailResult
func ReleaseItemDetailResult(v *ItemDetailResult) {
	v.ItemDetail = nil
	v.ItemInfo = nil
	poolItemDetailResult.Put(v)
}
