package nrt

import (
	"sync"
)

// StallSigningRespDto 结构体
type StallSigningRespDto struct {
	// 申请单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 摊位/门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolStallSigningRespDto = sync.Pool{
	New: func() any {
		return new(StallSigningRespDto)
	},
}

// GetStallSigningRespDto() 从对象池中获取StallSigningRespDto
func GetStallSigningRespDto() *StallSigningRespDto {
	return poolStallSigningRespDto.Get().(*StallSigningRespDto)
}

// ReleaseStallSigningRespDto 释放StallSigningRespDto
func ReleaseStallSigningRespDto(v *StallSigningRespDto) {
	v.OrderId = ""
	v.StoreId = 0
	poolStallSigningRespDto.Put(v)
}
