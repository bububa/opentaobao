package wdk

import (
	"sync"
)

// DeletePurchasePriceRequest 结构体
type DeletePurchasePriceRequest struct {
	// 请求幂等ID
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 门店ID
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 渠道
	ChannelCodes string `json:"channel_codes,omitempty" xml:"channel_codes,omitempty"`
}

var poolDeletePurchasePriceRequest = sync.Pool{
	New: func() any {
		return new(DeletePurchasePriceRequest)
	},
}

// GetDeletePurchasePriceRequest() 从对象池中获取DeletePurchasePriceRequest
func GetDeletePurchasePriceRequest() *DeletePurchasePriceRequest {
	return poolDeletePurchasePriceRequest.Get().(*DeletePurchasePriceRequest)
}

// ReleaseDeletePurchasePriceRequest 释放DeletePurchasePriceRequest
func ReleaseDeletePurchasePriceRequest(v *DeletePurchasePriceRequest) {
	v.OutId = ""
	v.SkuCode = ""
	v.OuCode = ""
	v.ChannelCodes = ""
	poolDeletePurchasePriceRequest.Put(v)
}
