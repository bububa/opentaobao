package waybill

import (
	"sync"
)

// WaybillServiceType 结构体
type WaybillServiceType struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolWaybillServiceType = sync.Pool{
	New: func() any {
		return new(WaybillServiceType)
	},
}

// GetWaybillServiceType() 从对象池中获取WaybillServiceType
func GetWaybillServiceType() *WaybillServiceType {
	return poolWaybillServiceType.Get().(*WaybillServiceType)
}

// ReleaseWaybillServiceType 释放WaybillServiceType
func ReleaseWaybillServiceType(v *WaybillServiceType) {
	v.Code = ""
	v.Name = ""
	poolWaybillServiceType.Put(v)
}
