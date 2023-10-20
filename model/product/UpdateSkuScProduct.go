package product

import (
	"sync"
)

// UpdateSkuScProduct 结构体
type UpdateSkuScProduct struct {
	// 商家外部id，如果填写，将以商家外部id查找被更新的SKU
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充；如果填写将以属性对形式查找被更新SKU
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 被更新的Sku对应的货品的货主id
	ScUserId int64 `json:"sc_user_id,omitempty" xml:"sc_user_id,omitempty"`
	// 被更新的Sku对应的货品的货品id
	ScProductId int64 `json:"sc_product_id,omitempty" xml:"sc_product_id,omitempty"`
	// SkuID，如果填写，将以SKUID查找被更新的SKU
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolUpdateSkuScProduct = sync.Pool{
	New: func() any {
		return new(UpdateSkuScProduct)
	},
}

// GetUpdateSkuScProduct() 从对象池中获取UpdateSkuScProduct
func GetUpdateSkuScProduct() *UpdateSkuScProduct {
	return poolUpdateSkuScProduct.Get().(*UpdateSkuScProduct)
}

// ReleaseUpdateSkuScProduct 释放UpdateSkuScProduct
func ReleaseUpdateSkuScProduct(v *UpdateSkuScProduct) {
	v.OuterId = ""
	v.Properties = ""
	v.ScUserId = 0
	v.ScProductId = 0
	v.SkuId = 0
	poolUpdateSkuScProduct.Put(v)
}
