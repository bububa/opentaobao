package trade

import (
	"sync"
)

// ChannelOrderItem 结构体
type ChannelOrderItem struct {
	// 条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品id（商品id和货号必填一个）
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货号（商品id和货号必填一个）
	InventoryNo string `json:"inventory_no,omitempty" xml:"inventory_no,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 分销价格(分)
	DistributionPrice int64 `json:"distribution_price,omitempty" xml:"distribution_price,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolChannelOrderItem = sync.Pool{
	New: func() any {
		return new(ChannelOrderItem)
	},
}

// GetChannelOrderItem() 从对象池中获取ChannelOrderItem
func GetChannelOrderItem() *ChannelOrderItem {
	return poolChannelOrderItem.Get().(*ChannelOrderItem)
}

// ReleaseChannelOrderItem 释放ChannelOrderItem
func ReleaseChannelOrderItem(v *ChannelOrderItem) {
	v.Barcode = ""
	v.SkuId = ""
	v.ItemId = ""
	v.InventoryNo = ""
	v.ItemName = ""
	v.DistributionPrice = 0
	v.Quantity = 0
	poolChannelOrderItem.Put(v)
}
