package tmallservice

import (
	"sync"
)

// WorkerQuitServiceStoreForTopReqDto 结构体
type WorkerQuitServiceStoreForTopReqDto struct {
	// 退出网点的工人身份证列表
	IdNumberList []string `json:"id_number_list,omitempty" xml:"id_number_list>string,omitempty"`
	// 网点编码
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
}

var poolWorkerQuitServiceStoreForTopReqDto = sync.Pool{
	New: func() any {
		return new(WorkerQuitServiceStoreForTopReqDto)
	},
}

// GetWorkerQuitServiceStoreForTopReqDto() 从对象池中获取WorkerQuitServiceStoreForTopReqDto
func GetWorkerQuitServiceStoreForTopReqDto() *WorkerQuitServiceStoreForTopReqDto {
	return poolWorkerQuitServiceStoreForTopReqDto.Get().(*WorkerQuitServiceStoreForTopReqDto)
}

// ReleaseWorkerQuitServiceStoreForTopReqDto 释放WorkerQuitServiceStoreForTopReqDto
func ReleaseWorkerQuitServiceStoreForTopReqDto(v *WorkerQuitServiceStoreForTopReqDto) {
	v.IdNumberList = v.IdNumberList[:0]
	v.ServiceStoreCode = ""
	poolWorkerQuitServiceStoreForTopReqDto.Put(v)
}
