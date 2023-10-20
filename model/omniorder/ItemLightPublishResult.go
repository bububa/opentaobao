package omniorder

import (
	"sync"
)

// ItemLightPublishResult 结构体
type ItemLightPublishResult struct {
	// 重复商品信息
	DuplicateInfos []ItemSkuDuplicateInfo `json:"duplicate_infos,omitempty" xml:"duplicate_infos>item_sku_duplicate_info,omitempty"`
	// 创建生成的skuId
	SkuIds string `json:"sku_ids,omitempty" xml:"sku_ids,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 创建生成的itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// data
	Data *ItemLightPublishResult `json:"data,omitempty" xml:"data,omitempty"`
}

var poolItemLightPublishResult = sync.Pool{
	New: func() any {
		return new(ItemLightPublishResult)
	},
}

// GetItemLightPublishResult() 从对象池中获取ItemLightPublishResult
func GetItemLightPublishResult() *ItemLightPublishResult {
	return poolItemLightPublishResult.Get().(*ItemLightPublishResult)
}

// ReleaseItemLightPublishResult 释放ItemLightPublishResult
func ReleaseItemLightPublishResult(v *ItemLightPublishResult) {
	v.DuplicateInfos = v.DuplicateInfos[:0]
	v.SkuIds = ""
	v.Code = ""
	v.Message = ""
	v.ItemId = 0
	v.Data = nil
	poolItemLightPublishResult.Put(v)
}
