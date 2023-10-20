package waybill

import (
	"sync"
)

// WaybillProductType 结构体
type WaybillProductType struct {
	// 物流服务
	ServiceTypes []WaybillServiceType `json:"service_types,omitempty" xml:"service_types>waybill_service_type,omitempty"`
	// 产品code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 产品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolWaybillProductType = sync.Pool{
	New: func() any {
		return new(WaybillProductType)
	},
}

// GetWaybillProductType() 从对象池中获取WaybillProductType
func GetWaybillProductType() *WaybillProductType {
	return poolWaybillProductType.Get().(*WaybillProductType)
}

// ReleaseWaybillProductType 释放WaybillProductType
func ReleaseWaybillProductType(v *WaybillProductType) {
	v.ServiceTypes = v.ServiceTypes[:0]
	v.Code = ""
	v.Name = ""
	poolWaybillProductType.Put(v)
}
