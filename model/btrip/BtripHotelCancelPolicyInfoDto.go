package btrip

import (
	"sync"
)

// BtripHotelCancelPolicyInfoDto 结构体
type BtripHotelCancelPolicyInfoDto struct {
	// 提前小时
	Hour int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// 罚金
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolBtripHotelCancelPolicyInfoDto = sync.Pool{
	New: func() any {
		return new(BtripHotelCancelPolicyInfoDto)
	},
}

// GetBtripHotelCancelPolicyInfoDto() 从对象池中获取BtripHotelCancelPolicyInfoDto
func GetBtripHotelCancelPolicyInfoDto() *BtripHotelCancelPolicyInfoDto {
	return poolBtripHotelCancelPolicyInfoDto.Get().(*BtripHotelCancelPolicyInfoDto)
}

// ReleaseBtripHotelCancelPolicyInfoDto 释放BtripHotelCancelPolicyInfoDto
func ReleaseBtripHotelCancelPolicyInfoDto(v *BtripHotelCancelPolicyInfoDto) {
	v.Hour = 0
	v.Value = 0
	poolBtripHotelCancelPolicyInfoDto.Put(v)
}
