package lsttrade

import (
	"sync"
)

// OrderBizInfo 结构体
type OrderBizInfo struct {
	// 仓库名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 零售通仓库类型。customer：虚仓；cainiao：实仓
	LstWarehouseType string `json:"lst_warehouse_type,omitempty" xml:"lst_warehouse_type,omitempty"`
}

var poolOrderBizInfo = sync.Pool{
	New: func() any {
		return new(OrderBizInfo)
	},
}

// GetOrderBizInfo() 从对象池中获取OrderBizInfo
func GetOrderBizInfo() *OrderBizInfo {
	return poolOrderBizInfo.Get().(*OrderBizInfo)
}

// ReleaseOrderBizInfo 释放OrderBizInfo
func ReleaseOrderBizInfo(v *OrderBizInfo) {
	v.WarehouseName = ""
	v.WarehouseCode = ""
	v.LstWarehouseType = ""
	poolOrderBizInfo.Put(v)
}
