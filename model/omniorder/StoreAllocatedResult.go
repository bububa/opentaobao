package omniorder

import (
	"sync"
)

// StoreAllocatedResult 结构体
type StoreAllocatedResult struct {
	// 0表示无系统异常
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 异常描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 店铺Id, 可能是门店或者电商仓
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 店铺类型, 门店(Store)或者电商仓(Warehouse)
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 店铺名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 订单分单状态, 可选值: Waiting(仍在派单中) Allocated(派单成功) AllocateFail(派单失败)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 子订单Id
	SubOid int64 `json:"sub_oid,omitempty" xml:"sub_oid,omitempty"`
	// 主订单Id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolStoreAllocatedResult = sync.Pool{
	New: func() any {
		return new(StoreAllocatedResult)
	},
}

// GetStoreAllocatedResult() 从对象池中获取StoreAllocatedResult
func GetStoreAllocatedResult() *StoreAllocatedResult {
	return poolStoreAllocatedResult.Get().(*StoreAllocatedResult)
}

// ReleaseStoreAllocatedResult 释放StoreAllocatedResult
func ReleaseStoreAllocatedResult(v *StoreAllocatedResult) {
	v.Code = ""
	v.Message = ""
	v.StoreId = ""
	v.StoreType = ""
	v.StoreName = ""
	v.Status = ""
	v.Attributes = ""
	v.SubOid = 0
	v.Tid = 0
	poolStoreAllocatedResult.Put(v)
}
