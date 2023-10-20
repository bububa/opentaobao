package tmallservice

import (
	"sync"
)

// SupplyCategoryWorkerDto 结构体
type SupplyCategoryWorkerDto struct {
	// 工人身份证号
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 工人手机号
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// 工人id
	WorkerId int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
}

var poolSupplyCategoryWorkerDto = sync.Pool{
	New: func() any {
		return new(SupplyCategoryWorkerDto)
	},
}

// GetSupplyCategoryWorkerDto() 从对象池中获取SupplyCategoryWorkerDto
func GetSupplyCategoryWorkerDto() *SupplyCategoryWorkerDto {
	return poolSupplyCategoryWorkerDto.Get().(*SupplyCategoryWorkerDto)
}

// ReleaseSupplyCategoryWorkerDto 释放SupplyCategoryWorkerDto
func ReleaseSupplyCategoryWorkerDto(v *SupplyCategoryWorkerDto) {
	v.IdNumber = ""
	v.WorkerMobile = ""
	v.WorkerId = 0
	poolSupplyCategoryWorkerDto.Put(v)
}
