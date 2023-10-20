package koubeimall

import (
	"sync"
)

// ItemServicesDetail 结构体
type ItemServicesDetail struct {
	// 服务明细列表
	ServiceDetails []string `json:"service_details,omitempty" xml:"service_details>string,omitempty"`
	// 服务名称
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

var poolItemServicesDetail = sync.Pool{
	New: func() any {
		return new(ItemServicesDetail)
	},
}

// GetItemServicesDetail() 从对象池中获取ItemServicesDetail
func GetItemServicesDetail() *ItemServicesDetail {
	return poolItemServicesDetail.Get().(*ItemServicesDetail)
}

// ReleaseItemServicesDetail 释放ItemServicesDetail
func ReleaseItemServicesDetail(v *ItemServicesDetail) {
	v.ServiceDetails = v.ServiceDetails[:0]
	v.ServiceName = ""
	poolItemServicesDetail.Put(v)
}
