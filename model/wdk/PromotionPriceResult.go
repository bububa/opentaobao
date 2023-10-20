package wdk

import (
	"sync"
)

// PromotionPriceResult 结构体
type PromotionPriceResult struct {
	// 促销信息记录
	ItemList []PromotionPriceDo `json:"item_list,omitempty" xml:"item_list>promotion_price_do,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 总数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 总页数
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 单页数据大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPromotionPriceResult = sync.Pool{
	New: func() any {
		return new(PromotionPriceResult)
	},
}

// GetPromotionPriceResult() 从对象池中获取PromotionPriceResult
func GetPromotionPriceResult() *PromotionPriceResult {
	return poolPromotionPriceResult.Get().(*PromotionPriceResult)
}

// ReleasePromotionPriceResult 释放PromotionPriceResult
func ReleasePromotionPriceResult(v *PromotionPriceResult) {
	v.ItemList = v.ItemList[:0]
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Total = 0
	v.PageCount = 0
	v.PageSize = 0
	v.PageIndex = 0
	v.Success = false
	poolPromotionPriceResult.Put(v)
}
