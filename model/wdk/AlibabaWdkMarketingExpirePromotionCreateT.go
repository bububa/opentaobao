package wdk

import (
	"sync"
)

// AlibabaWdkMarketingExpirePromotionCreateT 结构体
type AlibabaWdkMarketingExpirePromotionCreateT struct {
	// 商家code
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 商品itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkMarketingExpirePromotionCreateT = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingExpirePromotionCreateT)
	},
}

// GetAlibabaWdkMarketingExpirePromotionCreateT() 从对象池中获取AlibabaWdkMarketingExpirePromotionCreateT
func GetAlibabaWdkMarketingExpirePromotionCreateT() *AlibabaWdkMarketingExpirePromotionCreateT {
	return poolAlibabaWdkMarketingExpirePromotionCreateT.Get().(*AlibabaWdkMarketingExpirePromotionCreateT)
}

// ReleaseAlibabaWdkMarketingExpirePromotionCreateT 释放AlibabaWdkMarketingExpirePromotionCreateT
func ReleaseAlibabaWdkMarketingExpirePromotionCreateT(v *AlibabaWdkMarketingExpirePromotionCreateT) {
	v.MerchantCode = ""
	v.SkuCode = ""
	v.ShopId = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ItemId = 0
	v.Success = false
	poolAlibabaWdkMarketingExpirePromotionCreateT.Put(v)
}
