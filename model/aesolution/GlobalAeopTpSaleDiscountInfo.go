package aesolution

import (
	"sync"
)

// GlobalAeopTpSaleDiscountInfo 结构体
type GlobalAeopTpSaleDiscountInfo struct {
	// Promotion owner, can be Seller, Platform or Mix.
	PromotionOwner string `json:"promotion_owner,omitempty" xml:"promotion_owner,omitempty"`
	// promotion type
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// promotion ID
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// discount detail
	DiscountDetail *GlobalMoneyStr `json:"discount_detail,omitempty" xml:"discount_detail,omitempty"`
}

var poolGlobalAeopTpSaleDiscountInfo = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpSaleDiscountInfo)
	},
}

// GetGlobalAeopTpSaleDiscountInfo() 从对象池中获取GlobalAeopTpSaleDiscountInfo
func GetGlobalAeopTpSaleDiscountInfo() *GlobalAeopTpSaleDiscountInfo {
	return poolGlobalAeopTpSaleDiscountInfo.Get().(*GlobalAeopTpSaleDiscountInfo)
}

// ReleaseGlobalAeopTpSaleDiscountInfo 释放GlobalAeopTpSaleDiscountInfo
func ReleaseGlobalAeopTpSaleDiscountInfo(v *GlobalAeopTpSaleDiscountInfo) {
	v.PromotionOwner = ""
	v.PromotionType = ""
	v.PromotionId = ""
	v.DiscountDetail = nil
	poolGlobalAeopTpSaleDiscountInfo.Put(v)
}
