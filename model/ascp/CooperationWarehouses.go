package ascp

import (
	"sync"
)

// CooperationWarehouses 结构体
type CooperationWarehouses struct {
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 合作店铺仓code
	CooperationWhCode string `json:"cooperation_wh_code,omitempty" xml:"cooperation_wh_code,omitempty"`
}

var poolCooperationWarehouses = sync.Pool{
	New: func() any {
		return new(CooperationWarehouses)
	},
}

// GetCooperationWarehouses() 从对象池中获取CooperationWarehouses
func GetCooperationWarehouses() *CooperationWarehouses {
	return poolCooperationWarehouses.Get().(*CooperationWarehouses)
}

// ReleaseCooperationWarehouses 释放CooperationWarehouses
func ReleaseCooperationWarehouses(v *CooperationWarehouses) {
	v.WmsOwnerCode = ""
	v.CooperationWhCode = ""
	poolCooperationWarehouses.Put(v)
}
