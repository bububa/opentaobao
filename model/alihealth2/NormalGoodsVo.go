package alihealth2

import (
	"sync"
)

// NormalGoodsVo 结构体
type NormalGoodsVo struct {
	// 标品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 成本价，单位元
	CostPrice string `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
	// 售价，单位元
	SoldPrice string `json:"sold_price,omitempty" xml:"sold_price,omitempty"`
	// 标品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolNormalGoodsVo = sync.Pool{
	New: func() any {
		return new(NormalGoodsVo)
	},
}

// GetNormalGoodsVo() 从对象池中获取NormalGoodsVo
func GetNormalGoodsVo() *NormalGoodsVo {
	return poolNormalGoodsVo.Get().(*NormalGoodsVo)
}

// ReleaseNormalGoodsVo 释放NormalGoodsVo
func ReleaseNormalGoodsVo(v *NormalGoodsVo) {
	v.ItemName = ""
	v.CostPrice = ""
	v.SoldPrice = ""
	v.ItemId = 0
	poolNormalGoodsVo.Put(v)
}
