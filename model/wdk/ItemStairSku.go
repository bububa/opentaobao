package wdk

import (
	"sync"
)

// ItemStairSku 结构体
type ItemStairSku struct {
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 单位为分
	ExchangePrice int64 `json:"exchange_price,omitempty" xml:"exchange_price,omitempty"`
	// 换购商品总数限制
	ExchangeTotalLimit int64 `json:"exchange_total_limit,omitempty" xml:"exchange_total_limit,omitempty"`
	// 换购商品每单可换购数量限制
	ExchangeOrderLimit int64 `json:"exchange_order_limit,omitempty" xml:"exchange_order_limit,omitempty"`
}

var poolItemStairSku = sync.Pool{
	New: func() any {
		return new(ItemStairSku)
	},
}

// GetItemStairSku() 从对象池中获取ItemStairSku
func GetItemStairSku() *ItemStairSku {
	return poolItemStairSku.Get().(*ItemStairSku)
}

// ReleaseItemStairSku 释放ItemStairSku
func ReleaseItemStairSku(v *ItemStairSku) {
	v.SkuCode = ""
	v.ExchangePrice = 0
	v.ExchangeTotalLimit = 0
	v.ExchangeOrderLimit = 0
	poolItemStairSku.Put(v)
}
