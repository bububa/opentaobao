package wdk

import (
	"sync"
)

// ExpirePromotionResult 结构体
type ExpirePromotionResult struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// skuCode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// merchantCode
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// shopId
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolExpirePromotionResult = sync.Pool{
	New: func() any {
		return new(ExpirePromotionResult)
	},
}

// GetExpirePromotionResult() 从对象池中获取ExpirePromotionResult
func GetExpirePromotionResult() *ExpirePromotionResult {
	return poolExpirePromotionResult.Get().(*ExpirePromotionResult)
}

// ReleaseExpirePromotionResult 释放ExpirePromotionResult
func ReleaseExpirePromotionResult(v *ExpirePromotionResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.SkuCode = ""
	v.MerchantCode = ""
	v.ShopId = ""
	v.ItemId = 0
	v.Success = false
	poolExpirePromotionResult.Put(v)
}
