package ascpchannel

import (
	"sync"
)

// Transferunitorderitemdtos 结构体
type Transferunitorderitemdtos struct {
	// 品信息
	TransferExtendOrderItemList []Transferextendorderitemdtolist `json:"transfer_extend_order_item_list,omitempty" xml:"transfer_extend_order_item_list>transferextendorderitemdtolist,omitempty"`
	// LBX订单号
	UnitBizCode string `json:"unit_biz_code,omitempty" xml:"unit_biz_code,omitempty"`
	// 后端货品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 货品条码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 后端货品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 计划下发数量
	PlanQuantity int64 `json:"plan_quantity,omitempty" xml:"plan_quantity,omitempty"`
	// 实际回传数量
	RealQuantity int64 `json:"real_quantity,omitempty" xml:"real_quantity,omitempty"`
}

var poolTransferunitorderitemdtos = sync.Pool{
	New: func() any {
		return new(Transferunitorderitemdtos)
	},
}

// GetTransferunitorderitemdtos() 从对象池中获取Transferunitorderitemdtos
func GetTransferunitorderitemdtos() *Transferunitorderitemdtos {
	return poolTransferunitorderitemdtos.Get().(*Transferunitorderitemdtos)
}

// ReleaseTransferunitorderitemdtos 释放Transferunitorderitemdtos
func ReleaseTransferunitorderitemdtos(v *Transferunitorderitemdtos) {
	v.TransferExtendOrderItemList = v.TransferExtendOrderItemList[:0]
	v.UnitBizCode = ""
	v.ItemName = ""
	v.ItemCode = ""
	v.ItemId = 0
	v.PlanQuantity = 0
	v.RealQuantity = 0
	poolTransferunitorderitemdtos.Put(v)
}
