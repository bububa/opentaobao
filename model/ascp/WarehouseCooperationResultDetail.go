package ascp

import (
	"sync"
)

// WarehouseCooperationResultDetail 结构体
type WarehouseCooperationResultDetail struct {
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
}

var poolWarehouseCooperationResultDetail = sync.Pool{
	New: func() any {
		return new(WarehouseCooperationResultDetail)
	},
}

// GetWarehouseCooperationResultDetail() 从对象池中获取WarehouseCooperationResultDetail
func GetWarehouseCooperationResultDetail() *WarehouseCooperationResultDetail {
	return poolWarehouseCooperationResultDetail.Get().(*WarehouseCooperationResultDetail)
}

// ReleaseWarehouseCooperationResultDetail 释放WarehouseCooperationResultDetail
func ReleaseWarehouseCooperationResultDetail(v *WarehouseCooperationResultDetail) {
	v.WmsOwnerCode = ""
	poolWarehouseCooperationResultDetail.Put(v)
}
