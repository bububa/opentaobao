package wdk

import (
	"sync"
)

// MarketPageResult 结构体
type MarketPageResult struct {
	// 返回的数据
	ItemCouponSkuList []ItemCouponSku `json:"item_coupon_sku_list,omitempty" xml:"item_coupon_sku_list>item_coupon_sku,omitempty"`
	// 查询结果数据
	SkuList []ItemPoolSku `json:"sku_list,omitempty" xml:"sku_list>item_pool_sku,omitempty"`
	// 返回的数据
	ItemPoolSkuList []ItemPoolSku `json:"item_pool_sku_list,omitempty" xml:"item_pool_sku_list>item_pool_sku,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 参加当前活动的商品总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前分页
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 查询商品是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMarketPageResult = sync.Pool{
	New: func() any {
		return new(MarketPageResult)
	},
}

// GetMarketPageResult() 从对象池中获取MarketPageResult
func GetMarketPageResult() *MarketPageResult {
	return poolMarketPageResult.Get().(*MarketPageResult)
}

// ReleaseMarketPageResult 释放MarketPageResult
func ReleaseMarketPageResult(v *MarketPageResult) {
	v.ItemCouponSkuList = v.ItemCouponSkuList[:0]
	v.SkuList = v.SkuList[:0]
	v.ItemPoolSkuList = v.ItemPoolSkuList[:0]
	v.Message = ""
	v.ErrorCode = ""
	v.Total = 0
	v.PageSize = 0
	v.Current = 0
	v.TotalPage = 0
	v.Success = false
	poolMarketPageResult.Put(v)
}
