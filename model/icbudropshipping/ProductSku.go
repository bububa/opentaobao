package icbudropshipping

import (
	"sync"
)

// ProductSku 结构体
type ProductSku struct {
	// inventory list
	InventoryList []Inventory `json:"inventory_list,omitempty" xml:"inventory_list>inventory,omitempty"`
	// sku name value list
	SkuNameValueList []ProductSkuNameValue `json:"sku_name_value_list,omitempty" xml:"sku_name_value_list>product_sku_name_value,omitempty"`
	// ladder price list
	LadderPriceList []LadderPrice `json:"ladder_price_list,omitempty" xml:"ladder_price_list>ladder_price,omitempty"`
	// sku image url
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolProductSku = sync.Pool{
	New: func() any {
		return new(ProductSku)
	},
}

// GetProductSku() 从对象池中获取ProductSku
func GetProductSku() *ProductSku {
	return poolProductSku.Get().(*ProductSku)
}

// ReleaseProductSku 释放ProductSku
func ReleaseProductSku(v *ProductSku) {
	v.InventoryList = v.InventoryList[:0]
	v.SkuNameValueList = v.SkuNameValueList[:0]
	v.LadderPriceList = v.LadderPriceList[:0]
	v.ImageUrl = ""
	v.SkuId = 0
	poolProductSku.Put(v)
}
