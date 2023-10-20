package tbk

import (
	"sync"
)

// PromotionExtend 结构体
type PromotionExtend struct {
	// 权益推荐商品
	RecommendItemList []RecommendItemList `json:"recommend_item_list,omitempty" xml:"recommend_item_list>recommend_item_list,omitempty"`
	// 权益链接
	PromotionUrl string `json:"promotion_url,omitempty" xml:"promotion_url,omitempty"`
	// 有价券信息
	YoujiaCouponInfo *Youjiacouponinfo `json:"youjia_coupon_info,omitempty" xml:"youjia_coupon_info,omitempty"`
}

var poolPromotionExtend = sync.Pool{
	New: func() any {
		return new(PromotionExtend)
	},
}

// GetPromotionExtend() 从对象池中获取PromotionExtend
func GetPromotionExtend() *PromotionExtend {
	return poolPromotionExtend.Get().(*PromotionExtend)
}

// ReleasePromotionExtend 释放PromotionExtend
func ReleasePromotionExtend(v *PromotionExtend) {
	v.RecommendItemList = v.RecommendItemList[:0]
	v.PromotionUrl = ""
	v.YoujiaCouponInfo = nil
	poolPromotionExtend.Put(v)
}
