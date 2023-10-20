package tbrefund

import (
	"sync"
)

// CombineSubItemDo 结构体
type CombineSubItemDo struct {
	// 商品数字编号
	NumIid string `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品标题
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// sku_id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商家外部编码(可与商家外部系统对接)
	OuterIid string `json:"outer_iid,omitempty" xml:"outer_iid,omitempty"`
	// 123
	OuterSkuId string `json:"outer_sku_id,omitempty" xml:"outer_sku_id,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolCombineSubItemDo = sync.Pool{
	New: func() any {
		return new(CombineSubItemDo)
	},
}

// GetCombineSubItemDo() 从对象池中获取CombineSubItemDo
func GetCombineSubItemDo() *CombineSubItemDo {
	return poolCombineSubItemDo.Get().(*CombineSubItemDo)
}

// ReleaseCombineSubItemDo 释放CombineSubItemDo
func ReleaseCombineSubItemDo(v *CombineSubItemDo) {
	v.NumIid = ""
	v.ItemName = ""
	v.SkuId = ""
	v.OuterIid = ""
	v.OuterSkuId = ""
	v.Quantity = 0
	poolCombineSubItemDo.Put(v)
}
