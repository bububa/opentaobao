package wdk

import (
	"sync"
)

// DpsScanContainerMtopRequest 结构体
type DpsScanContainerMtopRequest struct {
	// 明细列表
	DetailIdList []int64 `json:"detail_id_list,omitempty" xml:"detail_id_list>int64,omitempty"`
	// 用户
	UserAccount string `json:"user_account,omitempty" xml:"user_account,omitempty"`
	// 容器号
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolDpsScanContainerMtopRequest = sync.Pool{
	New: func() any {
		return new(DpsScanContainerMtopRequest)
	},
}

// GetDpsScanContainerMtopRequest() 从对象池中获取DpsScanContainerMtopRequest
func GetDpsScanContainerMtopRequest() *DpsScanContainerMtopRequest {
	return poolDpsScanContainerMtopRequest.Get().(*DpsScanContainerMtopRequest)
}

// ReleaseDpsScanContainerMtopRequest 释放DpsScanContainerMtopRequest
func ReleaseDpsScanContainerMtopRequest(v *DpsScanContainerMtopRequest) {
	v.DetailIdList = v.DetailIdList[:0]
	v.UserAccount = ""
	v.ContainerCode = ""
	v.WarehouseCode = ""
	poolDpsScanContainerMtopRequest.Put(v)
}
