package tbitem

import (
	"sync"
)

// UpdateSkuQuantity 结构体
type UpdateSkuQuantity struct {
	// Sku的商家外部id，用于指定被修改库存的SKU
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充，用于指定被修改库存的SKU
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// SKU属于这个sku的商品的库存；增量编辑方式支持正数、负数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// SkuID，用于指定被修改库存的
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolUpdateSkuQuantity = sync.Pool{
	New: func() any {
		return new(UpdateSkuQuantity)
	},
}

// GetUpdateSkuQuantity() 从对象池中获取UpdateSkuQuantity
func GetUpdateSkuQuantity() *UpdateSkuQuantity {
	return poolUpdateSkuQuantity.Get().(*UpdateSkuQuantity)
}

// ReleaseUpdateSkuQuantity 释放UpdateSkuQuantity
func ReleaseUpdateSkuQuantity(v *UpdateSkuQuantity) {
	v.OuterId = ""
	v.Properties = ""
	v.Quantity = 0
	v.SkuId = 0
	poolUpdateSkuQuantity.Put(v)
}
