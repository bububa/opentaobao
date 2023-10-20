package omniorder

import (
	"sync"
)

// ItemSkuDuplicateInfo 结构体
type ItemSkuDuplicateInfo struct {
	// 重复商品详情列表，如果重复商品过多，目前只展示部分
	DuplicateDetails []ItemSkuDuplicateDetail `json:"duplicate_details,omitempty" xml:"duplicate_details>item_sku_duplicate_detail,omitempty"`
	// barcode
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// skuBarcode
	SkuBarcode string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
	// skuOuterId
	SkuOuterId string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
	// 重复商品数量
	DuplicateSize int64 `json:"duplicate_size,omitempty" xml:"duplicate_size,omitempty"`
}

var poolItemSkuDuplicateInfo = sync.Pool{
	New: func() any {
		return new(ItemSkuDuplicateInfo)
	},
}

// GetItemSkuDuplicateInfo() 从对象池中获取ItemSkuDuplicateInfo
func GetItemSkuDuplicateInfo() *ItemSkuDuplicateInfo {
	return poolItemSkuDuplicateInfo.Get().(*ItemSkuDuplicateInfo)
}

// ReleaseItemSkuDuplicateInfo 释放ItemSkuDuplicateInfo
func ReleaseItemSkuDuplicateInfo(v *ItemSkuDuplicateInfo) {
	v.DuplicateDetails = v.DuplicateDetails[:0]
	v.Barcode = ""
	v.OuterId = ""
	v.SkuBarcode = ""
	v.SkuOuterId = ""
	v.DuplicateSize = 0
	poolItemSkuDuplicateInfo.Put(v)
}
