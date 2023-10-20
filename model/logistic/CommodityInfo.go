package logistic

import (
	"sync"
)

// CommodityInfo 结构体
type CommodityInfo struct {
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品数量
	GoodsQuantity int64 `json:"goods_quantity,omitempty" xml:"goods_quantity,omitempty"`
}

var poolCommodityInfo = sync.Pool{
	New: func() any {
		return new(CommodityInfo)
	},
}

// GetCommodityInfo() 从对象池中获取CommodityInfo
func GetCommodityInfo() *CommodityInfo {
	return poolCommodityInfo.Get().(*CommodityInfo)
}

// ReleaseCommodityInfo 释放CommodityInfo
func ReleaseCommodityInfo(v *CommodityInfo) {
	v.ItemId = ""
	v.GoodsQuantity = 0
	poolCommodityInfo.Put(v)
}
