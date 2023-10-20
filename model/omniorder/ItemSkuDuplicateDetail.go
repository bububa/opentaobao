package omniorder

import (
	"sync"
)

// ItemSkuDuplicateDetail 结构体
type ItemSkuDuplicateDetail struct {
	// barcode
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// skuBarcode
	SkuBarcode string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
	// skuOuterId
	SkuOuterId string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolItemSkuDuplicateDetail = sync.Pool{
	New: func() any {
		return new(ItemSkuDuplicateDetail)
	},
}

// GetItemSkuDuplicateDetail() 从对象池中获取ItemSkuDuplicateDetail
func GetItemSkuDuplicateDetail() *ItemSkuDuplicateDetail {
	return poolItemSkuDuplicateDetail.Get().(*ItemSkuDuplicateDetail)
}

// ReleaseItemSkuDuplicateDetail 释放ItemSkuDuplicateDetail
func ReleaseItemSkuDuplicateDetail(v *ItemSkuDuplicateDetail) {
	v.Barcode = ""
	v.OuterId = ""
	v.SkuBarcode = ""
	v.SkuOuterId = ""
	v.ItemId = 0
	v.SkuId = 0
	poolItemSkuDuplicateDetail.Put(v)
}
