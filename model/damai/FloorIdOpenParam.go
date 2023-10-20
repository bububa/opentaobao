package damai

import (
	"sync"
)

// FloorIdOpenParam 结构体
type FloorIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 参数performId
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 参数systemId
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolFloorIdOpenParam = sync.Pool{
	New: func() any {
		return new(FloorIdOpenParam)
	},
}

// GetFloorIdOpenParam() 从对象池中获取FloorIdOpenParam
func GetFloorIdOpenParam() *FloorIdOpenParam {
	return poolFloorIdOpenParam.Get().(*FloorIdOpenParam)
}

// ReleaseFloorIdOpenParam 释放FloorIdOpenParam
func ReleaseFloorIdOpenParam(v *FloorIdOpenParam) {
	v.SupplierSecret = ""
	v.FloorId = 0
	v.PerformId = 0
	v.SystemId = 0
	poolFloorIdOpenParam.Put(v)
}
