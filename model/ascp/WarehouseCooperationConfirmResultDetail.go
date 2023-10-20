package ascp

import (
	"sync"
)

// WarehouseCooperationConfirmResultDetail 结构体
type WarehouseCooperationConfirmResultDetail struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 合作店铺仓code
	CooperationWhCode string `json:"cooperation_wh_code,omitempty" xml:"cooperation_wh_code,omitempty"`
}

var poolWarehouseCooperationConfirmResultDetail = sync.Pool{
	New: func() any {
		return new(WarehouseCooperationConfirmResultDetail)
	},
}

// GetWarehouseCooperationConfirmResultDetail() 从对象池中获取WarehouseCooperationConfirmResultDetail
func GetWarehouseCooperationConfirmResultDetail() *WarehouseCooperationConfirmResultDetail {
	return poolWarehouseCooperationConfirmResultDetail.Get().(*WarehouseCooperationConfirmResultDetail)
}

// ReleaseWarehouseCooperationConfirmResultDetail 释放WarehouseCooperationConfirmResultDetail
func ReleaseWarehouseCooperationConfirmResultDetail(v *WarehouseCooperationConfirmResultDetail) {
	v.Code = ""
	v.Message = ""
	v.WmsOwnerCode = ""
	v.CooperationWhCode = ""
	poolWarehouseCooperationConfirmResultDetail.Put(v)
}
