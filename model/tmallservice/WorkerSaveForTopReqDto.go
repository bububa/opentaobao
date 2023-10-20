package tmallservice

import (
	"sync"
)

// WorkerSaveForTopReqDto 结构体
type WorkerSaveForTopReqDto struct {
	// 身份证号码
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 用户地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 真实姓名
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 用户地址编码
	AddressId int64 `json:"address_id,omitempty" xml:"address_id,omitempty"`
	// 工人技能参数
	WorkerServiceAbility *WorkerServiceAbility `json:"worker_service_ability,omitempty" xml:"worker_service_ability,omitempty"`
	// 加入的网点参数
	JoinedStore *JoinedStore `json:"joined_store,omitempty" xml:"joined_store,omitempty"`
}

var poolWorkerSaveForTopReqDto = sync.Pool{
	New: func() any {
		return new(WorkerSaveForTopReqDto)
	},
}

// GetWorkerSaveForTopReqDto() 从对象池中获取WorkerSaveForTopReqDto
func GetWorkerSaveForTopReqDto() *WorkerSaveForTopReqDto {
	return poolWorkerSaveForTopReqDto.Get().(*WorkerSaveForTopReqDto)
}

// ReleaseWorkerSaveForTopReqDto 释放WorkerSaveForTopReqDto
func ReleaseWorkerSaveForTopReqDto(v *WorkerSaveForTopReqDto) {
	v.IdNumber = ""
	v.Address = ""
	v.RealName = ""
	v.Phone = ""
	v.AddressId = 0
	v.WorkerServiceAbility = nil
	v.JoinedStore = nil
	poolWorkerSaveForTopReqDto.Put(v)
}
