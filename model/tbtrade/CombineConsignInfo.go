package tbtrade

import (
	"sync"
)

// CombineConsignInfo 结构体
type CombineConsignInfo struct {
	// 套餐组合id
	CombId string `json:"comb_id,omitempty" xml:"comb_id,omitempty"`
	// 套餐内的商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 套餐内商品的skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 成分品的预计发货时间
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 成分品的预计发货时间
	RenderConsignTime string `json:"render_consign_time,omitempty" xml:"render_consign_time,omitempty"`
}

var poolCombineConsignInfo = sync.Pool{
	New: func() any {
		return new(CombineConsignInfo)
	},
}

// GetCombineConsignInfo() 从对象池中获取CombineConsignInfo
func GetCombineConsignInfo() *CombineConsignInfo {
	return poolCombineConsignInfo.Get().(*CombineConsignInfo)
}

// ReleaseCombineConsignInfo 释放CombineConsignInfo
func ReleaseCombineConsignInfo(v *CombineConsignInfo) {
	v.CombId = ""
	v.ItemId = ""
	v.SkuId = ""
	v.ConsignTime = ""
	v.RenderConsignTime = ""
	poolCombineConsignInfo.Put(v)
}
