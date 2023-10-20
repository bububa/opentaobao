package wdk

import (
	"sync"
)

// ItemDiscountActivitySkuQueryRequest 结构体
type ItemDiscountActivitySkuQueryRequest struct {
	// 商品编码列表
	SkuCodes []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
	// 商品条码列表
	BarCodes []string `json:"bar_codes,omitempty" xml:"bar_codes>string,omitempty"`
	// erp外部活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 分页查询参数
	PageInfo *ActivitySkuQueryDto `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// 换购品标识
	ExchangeSku bool `json:"exchange_sku,omitempty" xml:"exchange_sku,omitempty"`
}

var poolItemDiscountActivitySkuQueryRequest = sync.Pool{
	New: func() any {
		return new(ItemDiscountActivitySkuQueryRequest)
	},
}

// GetItemDiscountActivitySkuQueryRequest() 从对象池中获取ItemDiscountActivitySkuQueryRequest
func GetItemDiscountActivitySkuQueryRequest() *ItemDiscountActivitySkuQueryRequest {
	return poolItemDiscountActivitySkuQueryRequest.Get().(*ItemDiscountActivitySkuQueryRequest)
}

// ReleaseItemDiscountActivitySkuQueryRequest 释放ItemDiscountActivitySkuQueryRequest
func ReleaseItemDiscountActivitySkuQueryRequest(v *ItemDiscountActivitySkuQueryRequest) {
	v.SkuCodes = v.SkuCodes[:0]
	v.BarCodes = v.BarCodes[:0]
	v.OutActId = ""
	v.ActId = 0
	v.PageInfo = nil
	v.ExchangeSku = false
	poolItemDiscountActivitySkuQueryRequest.Put(v)
}
