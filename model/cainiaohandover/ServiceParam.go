package cainiaohandover

import (
	"sync"
)

// ServiceParam 结构体
type ServiceParam struct {
	// DOOR_PICKUP：揽收仓资源、SELF_SEND：自送dropOff
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolServiceParam = sync.Pool{
	New: func() any {
		return new(ServiceParam)
	},
}

// GetServiceParam() 从对象池中获取ServiceParam
func GetServiceParam() *ServiceParam {
	return poolServiceParam.Get().(*ServiceParam)
}

// ReleaseServiceParam 释放ServiceParam
func ReleaseServiceParam(v *ServiceParam) {
	v.Code = ""
	poolServiceParam.Put(v)
}
