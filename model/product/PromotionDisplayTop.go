package product

import (
	"sync"
)

// PromotionDisplayTop 结构体
type PromotionDisplayTop struct {
	// 单品级优惠信息
	PromotionInItem []PromotionInItem `json:"promotion_in_item,omitempty" xml:"promotion_in_item>promotion_in_item,omitempty"`
	// 店铺级优惠信息
	PromotionInShop []PromotionInShop `json:"promotion_in_shop,omitempty" xml:"promotion_in_shop>promotion_in_shop,omitempty"`
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
	v.PromotionInItem = v.PromotionInItem[:0]
	v.PromotionInShop = v.PromotionInShop[:0]
	poolPromotionDisplayTop.Put(v)
}
