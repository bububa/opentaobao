package wdk

import (
	"sync"
)

// BomItemInfos 结构体
type BomItemInfos struct {
	// quantity
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// itemCode
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// cabinetCode
	CabinetCode string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
}

var poolBomItemInfos = sync.Pool{
	New: func() any {
		return new(BomItemInfos)
	},
}

// GetBomItemInfos() 从对象池中获取BomItemInfos
func GetBomItemInfos() *BomItemInfos {
	return poolBomItemInfos.Get().(*BomItemInfos)
}

// ReleaseBomItemInfos 释放BomItemInfos
func ReleaseBomItemInfos(v *BomItemInfos) {
	v.Quantity = ""
	v.ItemCode = ""
	v.CabinetCode = ""
	poolBomItemInfos.Put(v)
}
