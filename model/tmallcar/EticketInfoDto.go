package tmallcar

import (
	"sync"
)

// EticketInfoDto 结构体
type EticketInfoDto struct {
	// 核销状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 核销状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolEticketInfoDto = sync.Pool{
	New: func() any {
		return new(EticketInfoDto)
	},
}

// GetEticketInfoDto() 从对象池中获取EticketInfoDto
func GetEticketInfoDto() *EticketInfoDto {
	return poolEticketInfoDto.Get().(*EticketInfoDto)
}

// ReleaseEticketInfoDto 释放EticketInfoDto
func ReleaseEticketInfoDto(v *EticketInfoDto) {
	v.Status = ""
	v.StatusDesc = ""
	v.OrderId = 0
	poolEticketInfoDto.Put(v)
}
