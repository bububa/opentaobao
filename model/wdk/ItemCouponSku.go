package wdk

import (
	"sync"
)

// ItemCouponSku 结构体
type ItemCouponSku struct {
	// 商品的skuCode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 淘宝item和shop的对应关系， k-itemId, v-shopId
	ItemShopRelation string `json:"item_shop_relation,omitempty" xml:"item_shop_relation,omitempty"`
}

var poolItemCouponSku = sync.Pool{
	New: func() any {
		return new(ItemCouponSku)
	},
}

// GetItemCouponSku() 从对象池中获取ItemCouponSku
func GetItemCouponSku() *ItemCouponSku {
	return poolItemCouponSku.Get().(*ItemCouponSku)
}

// ReleaseItemCouponSku 释放ItemCouponSku
func ReleaseItemCouponSku(v *ItemCouponSku) {
	v.SkuCode = ""
	v.SkuName = ""
	v.ItemShopRelation = ""
	poolItemCouponSku.Put(v)
}
