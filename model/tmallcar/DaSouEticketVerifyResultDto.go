package tmallcar

import (
	"sync"
)

// DaSouEticketVerifyResultDto 结构体
type DaSouEticketVerifyResultDto struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolDaSouEticketVerifyResultDto = sync.Pool{
	New: func() any {
		return new(DaSouEticketVerifyResultDto)
	},
}

// GetDaSouEticketVerifyResultDto() 从对象池中获取DaSouEticketVerifyResultDto
func GetDaSouEticketVerifyResultDto() *DaSouEticketVerifyResultDto {
	return poolDaSouEticketVerifyResultDto.Get().(*DaSouEticketVerifyResultDto)
}

// ReleaseDaSouEticketVerifyResultDto 释放DaSouEticketVerifyResultDto
func ReleaseDaSouEticketVerifyResultDto(v *DaSouEticketVerifyResultDto) {
	v.OrderId = 0
	poolDaSouEticketVerifyResultDto.Put(v)
}
