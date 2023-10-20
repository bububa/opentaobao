package tmallservice

import (
	"sync"
)

// WorkerRegisterForTopReqDto 结构体
type WorkerRegisterForTopReqDto struct {
	// 身份证
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 工人图像
	ProfilePictureUrl string `json:"profile_picture_url,omitempty" xml:"profile_picture_url,omitempty"`
	// 姓名
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 工人手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 详细地址编码
	AddressId int64 `json:"address_id,omitempty" xml:"address_id,omitempty"`
	// 工人能力model
	WorkerServiceAbility *WorkerServiceAbility `json:"worker_service_ability,omitempty" xml:"worker_service_ability,omitempty"`
	// 工人加入网点model
	JoinedStore *JoinedStore `json:"joined_store,omitempty" xml:"joined_store,omitempty"`
}

var poolWorkerRegisterForTopReqDto = sync.Pool{
	New: func() any {
		return new(WorkerRegisterForTopReqDto)
	},
}

// GetWorkerRegisterForTopReqDto() 从对象池中获取WorkerRegisterForTopReqDto
func GetWorkerRegisterForTopReqDto() *WorkerRegisterForTopReqDto {
	return poolWorkerRegisterForTopReqDto.Get().(*WorkerRegisterForTopReqDto)
}

// ReleaseWorkerRegisterForTopReqDto 释放WorkerRegisterForTopReqDto
func ReleaseWorkerRegisterForTopReqDto(v *WorkerRegisterForTopReqDto) {
	v.IdNumber = ""
	v.ProfilePictureUrl = ""
	v.RealName = ""
	v.AlipayAccount = ""
	v.Address = ""
	v.Phone = ""
	v.AddressId = 0
	v.WorkerServiceAbility = nil
	v.JoinedStore = nil
	poolWorkerRegisterForTopReqDto.Put(v)
}
