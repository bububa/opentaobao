package iot

import (
	"sync"
)

// PromotionDisplayTop 结构体
type PromotionDisplayTop struct {
	// 单品级优惠信息
	PromotionInItemList []PromotionInItem `json:"promotion_in_item_list,omitempty" xml:"promotion_in_item_list>promotion_in_item,omitempty"`
	// 店铺级优惠信息
	PromotionInShopList []PromotionInShop `json:"promotion_in_shop_list,omitempty" xml:"promotion_in_shop_list>promotion_in_shop,omitempty"`
}

var poolPromotionDisplayTop = sync.Pool{
	New: func() any {
		return new(PromotionDisplayTop)
	},
}

// GetPromotionDisplayTop() 从对象池中获取PromotionDisplayTop
func GetPromotionDisplayTop() *PromotionDisplayTop {
	return poolPromotionDisplayTop.Get().(*PromotionDisplayTop)
}

// ReleasePromotionDisplayTop 释放PromotionDisplayTop
func ReleasePromotionDisplayTop(v *PromotionDisplayTop) {
	v.PromotionInItemList = v.PromotionInItemList[:0]
	v.PromotionInShopList = v.PromotionInShopList[:0]
	poolPromotionDisplayTop.Put(v)
}
