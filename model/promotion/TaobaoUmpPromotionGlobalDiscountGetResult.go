package promotion

import (
	"sync"
)

// TaobaoUmpPromotionGlobalDiscountGetResult 结构体
type TaobaoUmpPromotionGlobalDiscountGetResult struct {
	// defaultModel
	DefaultModel *SellerGlobalDiscount `json:"default_model,omitempty" xml:"default_model,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoUmpPromotionGlobalDiscountGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoUmpPromotionGlobalDiscountGetResult)
	},
}

// GetTaobaoUmpPromotionGlobalDiscountGetResult() 从对象池中获取TaobaoUmpPromotionGlobalDiscountGetResult
func GetTaobaoUmpPromotionGlobalDiscountGetResult() *TaobaoUmpPromotionGlobalDiscountGetResult {
	return poolTaobaoUmpPromotionGlobalDiscountGetResult.Get().(*TaobaoUmpPromotionGlobalDiscountGetResult)
}

// ReleaseTaobaoUmpPromotionGlobalDiscountGetResult 释放TaobaoUmpPromotionGlobalDiscountGetResult
func ReleaseTaobaoUmpPromotionGlobalDiscountGetResult(v *TaobaoUmpPromotionGlobalDiscountGetResult) {
	v.DefaultModel = nil
	v.Success = false
	poolTaobaoUmpPromotionGlobalDiscountGetResult.Put(v)
}
