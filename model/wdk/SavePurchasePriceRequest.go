package wdk

import (
	"sync"
)

// SavePurchasePriceRequest 结构体
type SavePurchasePriceRequest struct {
	// 门店ID
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 幂等ID
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 渠道
	ChannelCodes string `json:"channel_codes,omitempty" xml:"channel_codes,omitempty"`
	// 2-经销，3-代销，6-寄售，默认为【6-寄售】
	MarketingType int64 `json:"marketing_type,omitempty" xml:"marketing_type,omitempty"`
	// 含税采购价格，单位【分】
	PurchasePriceWithTax int64 `json:"purchase_price_with_tax,omitempty" xml:"purchase_price_with_tax,omitempty"`
	// 1-基准价格，3-区间价格
	PriceType int64 `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 区间价格生效时间
	EffectiveStartTime int64 `json:"effective_start_time,omitempty" xml:"effective_start_time,omitempty"`
	// 区间价格失效时间
	EffectiveEndTime int64 `json:"effective_end_time,omitempty" xml:"effective_end_time,omitempty"`
}

var poolSavePurchasePriceRequest = sync.Pool{
	New: func() any {
		return new(SavePurchasePriceRequest)
	},
}

// GetSavePurchasePriceRequest() 从对象池中获取SavePurchasePriceRequest
func GetSavePurchasePriceRequest() *SavePurchasePriceRequest {
	return poolSavePurchasePriceRequest.Get().(*SavePurchasePriceRequest)
}

// ReleaseSavePurchasePriceRequest 释放SavePurchasePriceRequest
func ReleaseSavePurchasePriceRequest(v *SavePurchasePriceRequest) {
	v.OuCode = ""
	v.SkuCode = ""
	v.OutId = ""
	v.ChannelCodes = ""
	v.MarketingType = 0
	v.PurchasePriceWithTax = 0
	v.PriceType = 0
	v.EffectiveStartTime = 0
	v.EffectiveEndTime = 0
	poolSavePurchasePriceRequest.Put(v)
}
