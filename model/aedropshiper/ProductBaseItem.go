package aedropshiper

import (
	"sync"
)

// ProductBaseItem 结构体
type ProductBaseItem struct {
	// 商品sku
	SkuAttr string `json:"sku_attr,omitempty" xml:"sku_attr,omitempty"`
	// 物流服务名称
	LogisticsServiceName string `json:"logistics_service_name,omitempty" xml:"logistics_service_name,omitempty"`
	// 用户留言
	OrderMemo string `json:"order_memo,omitempty" xml:"order_memo,omitempty"`
	// 商品数量
	ProductCount int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolProductBaseItem = sync.Pool{
	New: func() any {
		return new(ProductBaseItem)
	},
}

// GetProductBaseItem() 从对象池中获取ProductBaseItem
func GetProductBaseItem() *ProductBaseItem {
	return poolProductBaseItem.Get().(*ProductBaseItem)
}

// ReleaseProductBaseItem 释放ProductBaseItem
func ReleaseProductBaseItem(v *ProductBaseItem) {
	v.SkuAttr = ""
	v.LogisticsServiceName = ""
	v.OrderMemo = ""
	v.ProductCount = 0
	v.ProductId = 0
	poolProductBaseItem.Put(v)
}
