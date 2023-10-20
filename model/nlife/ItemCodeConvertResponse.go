package nlife

import (
	"sync"
)

// ItemCodeConvertResponse 结构体
type ItemCodeConvertResponse struct {
	// 转码后的结果
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// BARCODE / UNIQUECODE
	CodeType string `json:"code_type,omitempty" xml:"code_type,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// storeId
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolItemCodeConvertResponse = sync.Pool{
	New: func() any {
		return new(ItemCodeConvertResponse)
	},
}

// GetItemCodeConvertResponse() 从对象池中获取ItemCodeConvertResponse
func GetItemCodeConvertResponse() *ItemCodeConvertResponse {
	return poolItemCodeConvertResponse.Get().(*ItemCodeConvertResponse)
}

// ReleaseItemCodeConvertResponse 释放ItemCodeConvertResponse
func ReleaseItemCodeConvertResponse(v *ItemCodeConvertResponse) {
	v.Code = ""
	v.CodeType = ""
	v.ItemId = 0
	v.SkuId = 0
	v.StoreId = 0
	poolItemCodeConvertResponse.Put(v)
}
