package ascpchannel

import (
	"sync"
)

// Inventorymainoperation 结构体
type Inventorymainoperation struct {
	// 子单操作明细
	DetailOperationList []Detailoperationlist `json:"detail_operation_list,omitempty" xml:"detail_operation_list>detailoperationlist,omitempty"`
	// 操作主单
	MainOrder *Mainorder `json:"main_order,omitempty" xml:"main_order,omitempty"`
}

var poolInventorymainoperation = sync.Pool{
	New: func() any {
		return new(Inventorymainoperation)
	},
}

// GetInventorymainoperation() 从对象池中获取Inventorymainoperation
func GetInventorymainoperation() *Inventorymainoperation {
	return poolInventorymainoperation.Get().(*Inventorymainoperation)
}

// ReleaseInventorymainoperation 释放Inventorymainoperation
func ReleaseInventorymainoperation(v *Inventorymainoperation) {
	v.DetailOperationList = v.DetailOperationList[:0]
	v.MainOrder = nil
	poolInventorymainoperation.Put(v)
}
