package ascp

import (
	"sync"
)

// ScItemInfo 结构体
type ScItemInfo struct {
	// 货品id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 销售库存总数
	SaleTotalCount int64 `json:"sale_total_count,omitempty" xml:"sale_total_count,omitempty"`
	// 销售库存占用
	SaleOccupyCount int64 `json:"sale_occupy_count,omitempty" xml:"sale_occupy_count,omitempty"`
	// 销售库存可用
	SaleRemainCount int64 `json:"sale_remain_count,omitempty" xml:"sale_remain_count,omitempty"`
}

var poolScItemInfo = sync.Pool{
	New: func() any {
		return new(ScItemInfo)
	},
}

// GetScItemInfo() 从对象池中获取ScItemInfo
func GetScItemInfo() *ScItemInfo {
	return poolScItemInfo.Get().(*ScItemInfo)
}

// ReleaseScItemInfo 释放ScItemInfo
func ReleaseScItemInfo(v *ScItemInfo) {
	v.ScItemId = ""
	v.ScItemCode = ""
	v.WarehouseCode = ""
	v.SaleTotalCount = 0
	v.SaleOccupyCount = 0
	v.SaleRemainCount = 0
	poolScItemInfo.Put(v)
}
