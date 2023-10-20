package wdk

import (
	"sync"
)

// PromotionContentResult 结构体
type PromotionContentResult struct {
	// 错误信息
	ErrorMessageList []string `json:"error_message_list,omitempty" xml:"error_message_list>string,omitempty"`
	// 唯一标示
	SkuCodeUniqueIdStr string `json:"sku_code_unique_id_str,omitempty" xml:"sku_code_unique_id_str,omitempty"`
}

var poolPromotionContentResult = sync.Pool{
	New: func() any {
		return new(PromotionContentResult)
	},
}

// GetPromotionContentResult() 从对象池中获取PromotionContentResult
func GetPromotionContentResult() *PromotionContentResult {
	return poolPromotionContentResult.Get().(*PromotionContentResult)
}

// ReleasePromotionContentResult 释放PromotionContentResult
func ReleasePromotionContentResult(v *PromotionContentResult) {
	v.ErrorMessageList = v.ErrorMessageList[:0]
	v.SkuCodeUniqueIdStr = ""
	poolPromotionContentResult.Put(v)
}
