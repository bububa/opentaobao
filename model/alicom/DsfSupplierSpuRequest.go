package alicom

import (
	"sync"
)

// DsfSupplierSpuRequest 结构体
type DsfSupplierSpuRequest struct {
	// 分组
	GroupList []GroupRequest `json:"group_list,omitempty" xml:"group_list>group_request,omitempty"`
	// 业务类型
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
}

var poolDsfSupplierSpuRequest = sync.Pool{
	New: func() any {
		return new(DsfSupplierSpuRequest)
	},
}

// GetDsfSupplierSpuRequest() 从对象池中获取DsfSupplierSpuRequest
func GetDsfSupplierSpuRequest() *DsfSupplierSpuRequest {
	return poolDsfSupplierSpuRequest.Get().(*DsfSupplierSpuRequest)
}

// ReleaseDsfSupplierSpuRequest 释放DsfSupplierSpuRequest
func ReleaseDsfSupplierSpuRequest(v *DsfSupplierSpuRequest) {
	v.GroupList = v.GroupList[:0]
	v.BusinessType = ""
	poolDsfSupplierSpuRequest.Put(v)
}
