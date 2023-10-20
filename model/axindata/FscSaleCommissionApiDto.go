package axindata

import (
	"sync"
)

// FscSaleCommissionApiDto 结构体
type FscSaleCommissionApiDto struct {
	// 佣金类型
	CommissionType string `json:"commission_type,omitempty" xml:"commission_type,omitempty"`
	// 佣金数值
	CommissionNum string `json:"commission_num,omitempty" xml:"commission_num,omitempty"`
}

var poolFscSaleCommissionApiDto = sync.Pool{
	New: func() any {
		return new(FscSaleCommissionApiDto)
	},
}

// GetFscSaleCommissionApiDto() 从对象池中获取FscSaleCommissionApiDto
func GetFscSaleCommissionApiDto() *FscSaleCommissionApiDto {
	return poolFscSaleCommissionApiDto.Get().(*FscSaleCommissionApiDto)
}

// ReleaseFscSaleCommissionApiDto 释放FscSaleCommissionApiDto
func ReleaseFscSaleCommissionApiDto(v *FscSaleCommissionApiDto) {
	v.CommissionType = ""
	v.CommissionNum = ""
	poolFscSaleCommissionApiDto.Put(v)
}
