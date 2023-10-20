package wdk

import (
	"sync"
)

// WarehouseLossReason 结构体
type WarehouseLossReason struct {
	// 报损原因描述
	WarehouseLossReason string `json:"warehouse_loss_reason,omitempty" xml:"warehouse_loss_reason,omitempty"`
	// 报损原因编码
	WarehouseLossReasonCode string `json:"warehouse_loss_reason_code,omitempty" xml:"warehouse_loss_reason_code,omitempty"`
	// 仓内报损数量
	WarehouseLossStockQuantity string `json:"warehouse_loss_stock_quantity,omitempty" xml:"warehouse_loss_stock_quantity,omitempty"`
}

var poolWarehouseLossReason = sync.Pool{
	New: func() any {
		return new(WarehouseLossReason)
	},
}

// GetWarehouseLossReason() 从对象池中获取WarehouseLossReason
func GetWarehouseLossReason() *WarehouseLossReason {
	return poolWarehouseLossReason.Get().(*WarehouseLossReason)
}

// ReleaseWarehouseLossReason 释放WarehouseLossReason
func ReleaseWarehouseLossReason(v *WarehouseLossReason) {
	v.WarehouseLossReason = ""
	v.WarehouseLossReasonCode = ""
	v.WarehouseLossStockQuantity = ""
	poolWarehouseLossReason.Put(v)
}
