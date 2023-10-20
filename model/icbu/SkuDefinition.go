package icbu

import (
	"sync"
)

// SkuDefinition 结构体
type SkuDefinition struct {
	// 根据订单数量设置折扣价
	BulkDiscountPrices []BulkDiscountPrice `json:"bulk_discount_prices,omitempty" xml:"bulk_discount_prices>bulk_discount_price,omitempty"`
	// 商品的库存列表
	InventoryDtoList []ProductInventoryDto `json:"inventory_dto_list,omitempty" xml:"inventory_dto_list>product_inventory_dto,omitempty"`
	// attr2Value
	Attr2Value string `json:"attr2_value,omitempty" xml:"attr2_value,omitempty"`
	// 商品的SKU编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品的SKUid，唯一标识SKU
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolSkuDefinition = sync.Pool{
	New: func() any {
		return new(SkuDefinition)
	},
}

// GetSkuDefinition() 从对象池中获取SkuDefinition
func GetSkuDefinition() *SkuDefinition {
	return poolSkuDefinition.Get().(*SkuDefinition)
}

// ReleaseSkuDefinition 释放SkuDefinition
func ReleaseSkuDefinition(v *SkuDefinition) {
	v.BulkDiscountPrices = v.BulkDiscountPrices[:0]
	v.InventoryDtoList = v.InventoryDtoList[:0]
	v.Attr2Value = ""
	v.SkuCode = ""
	v.SkuId = 0
	poolSkuDefinition.Put(v)
}
