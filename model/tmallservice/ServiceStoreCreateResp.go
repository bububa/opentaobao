package tmallservice

import (
	"sync"
)

// ServiceStoreCreateResp 结构体
type ServiceStoreCreateResp struct {
	// 门店密钥
	Secret string `json:"secret,omitempty" xml:"secret,omitempty"`
	// 门店id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolServiceStoreCreateResp = sync.Pool{
	New: func() any {
		return new(ServiceStoreCreateResp)
	},
}

// GetServiceStoreCreateResp() 从对象池中获取ServiceStoreCreateResp
func GetServiceStoreCreateResp() *ServiceStoreCreateResp {
	return poolServiceStoreCreateResp.Get().(*ServiceStoreCreateResp)
}

// ReleaseServiceStoreCreateResp 释放ServiceStoreCreateResp
func ReleaseServiceStoreCreateResp(v *ServiceStoreCreateResp) {
	v.Secret = ""
	v.Id = 0
	poolServiceStoreCreateResp.Put(v)
}
