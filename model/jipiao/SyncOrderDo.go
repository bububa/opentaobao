package jipiao

import (
	"sync"
)

// SyncOrderDo 结构体
type SyncOrderDo struct {
	// 改签后的舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 改签前的舱位
	LastCabin string `json:"last_cabin,omitempty" xml:"last_cabin,omitempty"`
	// 改签备注信息
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 改签后航班信息
	ModifyAfterFlight *Flight `json:"modify_after_flight,omitempty" xml:"modify_after_flight,omitempty"`
	// 改签前航班信息
	ModifyBeforeFlight *Flight `json:"modify_before_flight,omitempty" xml:"modify_before_flight,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 乘客信息
	Passenger *Passenger `json:"passenger,omitempty" xml:"passenger,omitempty"`
	// 改签费(单位分)
	ModifyFee int64 `json:"modify_fee,omitempty" xml:"modify_fee,omitempty"`
	// 升舱费(单位分)
	UpgradeFee int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// 申请单状态。1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSyncOrderDo = sync.Pool{
	New: func() any {
		return new(SyncOrderDo)
	},
}

// GetSyncOrderDo() 从对象池中获取SyncOrderDo
func GetSyncOrderDo() *SyncOrderDo {
	return poolSyncOrderDo.Get().(*SyncOrderDo)
}

// ReleaseSyncOrderDo 释放SyncOrderDo
func ReleaseSyncOrderDo(v *SyncOrderDo) {
	v.Cabin = ""
	v.LastCabin = ""
	v.Memo = ""
	v.ApplyId = 0
	v.ModifyAfterFlight = nil
	v.ModifyBeforeFlight = nil
	v.OrderId = 0
	v.Passenger = nil
	v.ModifyFee = 0
	v.UpgradeFee = 0
	v.Status = 0
	poolSyncOrderDo.Put(v)
}
