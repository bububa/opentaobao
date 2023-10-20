package waybill

import (
	"sync"
)

// ReachableRecommendRequestDto 结构体
type ReachableRecommendRequestDto struct {
	// 指定快递公司
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 地址
	Address *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
}

var poolReachableRecommendRequestDto = sync.Pool{
	New: func() any {
		return new(ReachableRecommendRequestDto)
	},
}

// GetReachableRecommendRequestDto() 从对象池中获取ReachableRecommendRequestDto
func GetReachableRecommendRequestDto() *ReachableRecommendRequestDto {
	return poolReachableRecommendRequestDto.Get().(*ReachableRecommendRequestDto)
}

// ReleaseReachableRecommendRequestDto 释放ReachableRecommendRequestDto
func ReleaseReachableRecommendRequestDto(v *ReachableRecommendRequestDto) {
	v.CpCode = ""
	v.Address = nil
	poolReachableRecommendRequestDto.Put(v)
}
