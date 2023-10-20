package omniorder

import (
	"sync"
)

// QuantityDetail 结构体
type QuantityDetail struct {
	// 库存类型
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 占用库存
	OccupyQuantity int64 `json:"occupy_quantity,omitempty" xml:"occupy_quantity,omitempty"`
}

var poolQuantityDetail = sync.Pool{
	New: func() any {
		return new(QuantityDetail)
	},
}

// GetQuantityDetail() 从对象池中获取QuantityDetail
func GetQuantityDetail() *QuantityDetail {
	return poolQuantityDetail.Get().(*QuantityDetail)
}

// ReleaseQuantityDetail 释放QuantityDetail
func ReleaseQuantityDetail(v *QuantityDetail) {
	v.InventoryType = ""
	v.Quantity = 0
	v.OccupyQuantity = 0
	poolQuantityDetail.Put(v)
}
