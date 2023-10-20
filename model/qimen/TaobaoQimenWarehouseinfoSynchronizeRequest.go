package qimen

import (
	"sync"
)

// TaobaoQimenWarehouseinfoSynchronizeRequest 结构体
type TaobaoQimenWarehouseinfoSynchronizeRequest struct {
	// 仓库信息
	WarehouseInfos []WarehouseInfos `json:"warehouseInfos,omitempty" xml:"warehouseInfos>warehouse_infos,omitempty"`
	// 货主编码，string（50）
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 货主名称，string（50）
	OwnerName string `json:"ownerName,omitempty" xml:"ownerName,omitempty"`
}

var poolTaobaoQimenWarehouseinfoSynchronizeRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWarehouseinfoSynchronizeRequest)
	},
}

// GetTaobaoQimenWarehouseinfoSynchronizeRequest() 从对象池中获取TaobaoQimenWarehouseinfoSynchronizeRequest
func GetTaobaoQimenWarehouseinfoSynchronizeRequest() *TaobaoQimenWarehouseinfoSynchronizeRequest {
	return poolTaobaoQimenWarehouseinfoSynchronizeRequest.Get().(*TaobaoQimenWarehouseinfoSynchronizeRequest)
}

// ReleaseTaobaoQimenWarehouseinfoSynchronizeRequest 释放TaobaoQimenWarehouseinfoSynchronizeRequest
func ReleaseTaobaoQimenWarehouseinfoSynchronizeRequest(v *TaobaoQimenWarehouseinfoSynchronizeRequest) {
	v.WarehouseInfos = v.WarehouseInfos[:0]
	v.OwnerCode = ""
	v.OwnerName = ""
	poolTaobaoQimenWarehouseinfoSynchronizeRequest.Put(v)
}
