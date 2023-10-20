package ascpchannel

import (
	"sync"
)

// Transferdetaildtolist 结构体
type Transferdetaildtolist struct {
	// 品基本信息
	TransferUnitOrderItemList []Transferunitorderitemdtos `json:"transfer_unit_order_item_list,omitempty" xml:"transfer_unit_order_item_list>transferunitorderitemdtos,omitempty"`
	// 仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// LBX订单号
	UnitBizCode string `json:"unit_biz_code,omitempty" xml:"unit_biz_code,omitempty"`
	// 1-出库单；2-入库单
	UnitType string `json:"unit_type,omitempty" xml:"unit_type,omitempty"`
	// 出库lbx-下发仓、仓接单、部分出、全出 	 * 入库lbx-下发仓、仓接单、部分入、全入
	FulfilUniBizStatus string `json:"fulfil_uni_biz_status,omitempty" xml:"fulfil_uni_biz_status,omitempty"`
}

var poolTransferdetaildtolist = sync.Pool{
	New: func() any {
		return new(Transferdetaildtolist)
	},
}

// GetTransferdetaildtolist() 从对象池中获取Transferdetaildtolist
func GetTransferdetaildtolist() *Transferdetaildtolist {
	return poolTransferdetaildtolist.Get().(*Transferdetaildtolist)
}

// ReleaseTransferdetaildtolist 释放Transferdetaildtolist
func ReleaseTransferdetaildtolist(v *Transferdetaildtolist) {
	v.TransferUnitOrderItemList = v.TransferUnitOrderItemList[:0]
	v.StoreCode = ""
	v.UnitBizCode = ""
	v.UnitType = ""
	v.FulfilUniBizStatus = ""
	poolTransferdetaildtolist.Put(v)
}
