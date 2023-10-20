package wdk

import (
	"sync"
)

// PromotionContent 结构体
type PromotionContent struct {
	// sku列表
	PromotionSkuList []PromotionSku `json:"promotion_sku_list,omitempty" xml:"promotion_sku_list>promotion_sku,omitempty"`
	// 客户商家编码
	CustomerMerchantCode string `json:"customer_merchant_code,omitempty" xml:"customer_merchant_code,omitempty"`
	// 客户编码
	CustomerCode string `json:"customer_code,omitempty" xml:"customer_code,omitempty"`
	// 客户门店
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 促销档期编码
	OuterPromotionCode string `json:"outer_promotion_code,omitempty" xml:"outer_promotion_code,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 进售价类型
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
}

var poolPromotionContent = sync.Pool{
	New: func() any {
		return new(PromotionContent)
	},
}

// GetPromotionContent() 从对象池中获取PromotionContent
func GetPromotionContent() *PromotionContent {
	return poolPromotionContent.Get().(*PromotionContent)
}

// ReleasePromotionContent 释放PromotionContent
func ReleasePromotionContent(v *PromotionContent) {
	v.PromotionSkuList = v.PromotionSkuList[:0]
	v.CustomerMerchantCode = ""
	v.CustomerCode = ""
	v.OuCode = ""
	v.OuterPromotionCode = ""
	v.MerchantCode = ""
	v.PromotionType = ""
	poolPromotionContent.Put(v)
}
