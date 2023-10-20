package ascp

import (
	"sync"
)

// WarehouseScItemRelation 结构体
type WarehouseScItemRelation struct {
	// 仓库编码
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
	// 仓库货品编码
	WarehouseScItemCode string `json:"warehouse_sc_item_code,omitempty" xml:"warehouse_sc_item_code,omitempty"`
}

var poolWarehouseScItemRelation = sync.Pool{
	New: func() any {
		return new(WarehouseScItemRelation)
	},
}

// GetWarehouseScItemRelation() 从对象池中获取WarehouseScItemRelation
func GetWarehouseScItemRelation() *WarehouseScItemRelation {
	return poolWarehouseScItemRelation.Get().(*WarehouseScItemRelation)
}

// ReleaseWarehouseScItemRelation 释放WarehouseScItemRelation
func ReleaseWarehouseScItemRelation(v *WarehouseScItemRelation) {
	v.ErpWarehouseCode = ""
	v.WarehouseScItemCode = ""
	poolWarehouseScItemRelation.Put(v)
}
