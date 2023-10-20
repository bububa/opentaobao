package travel

import (
	"sync"
)

// TopTravelItem 结构体
type TopTravelItem struct {
	// 商家自定义商品编码
	OutProductId string `json:"out_product_id,omitempty" xml:"out_product_id,omitempty"`
	// 商品修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 扩展信息
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 商品创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolTopTravelItem = sync.Pool{
	New: func() any {
		return new(TopTravelItem)
	},
}

// GetTopTravelItem() 从对象池中获取TopTravelItem
func GetTopTravelItem() *TopTravelItem {
	return poolTopTravelItem.Get().(*TopTravelItem)
}

// ReleaseTopTravelItem 释放TopTravelItem
func ReleaseTopTravelItem(v *TopTravelItem) {
	v.OutProductId = ""
	v.Modified = ""
	v.Extend = ""
	v.Created = ""
	v.OuterId = ""
	v.ItemId = 0
	v.SkuId = 0
	poolTopTravelItem.Put(v)
}
