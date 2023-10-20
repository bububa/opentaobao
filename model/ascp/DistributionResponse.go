package ascp

import (
	"sync"
)

// DistributionResponse 结构体
type DistributionResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误消息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 要选择进行铺货的店铺宝贝 itemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 要选择进行铺货的店铺宝贝 skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	//  处理是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDistributionResponse = sync.Pool{
	New: func() any {
		return new(DistributionResponse)
	},
}

// GetDistributionResponse() 从对象池中获取DistributionResponse
func GetDistributionResponse() *DistributionResponse {
	return poolDistributionResponse.Get().(*DistributionResponse)
}

// ReleaseDistributionResponse 释放DistributionResponse
func ReleaseDistributionResponse(v *DistributionResponse) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.ItemId = ""
	v.SkuId = ""
	v.ScItemId = 0
	v.Success = false
	poolDistributionResponse.Put(v)
}
