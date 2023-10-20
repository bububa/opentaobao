package util

import (
	"sync"
)

// PosInfoDto 结构体
type PosInfoDto struct {
	// 是否支持小数
	Support4Decimal string `json:"support4_decimal,omitempty" xml:"support4_decimal,omitempty"`
	// 专柜号
	CounterNo string `json:"counter_no,omitempty" xml:"counter_no,omitempty"`
	// 门店号
	StoreNo string `json:"store_no,omitempty" xml:"store_no,omitempty"`
}

var poolPosInfoDto = sync.Pool{
	New: func() any {
		return new(PosInfoDto)
	},
}

// GetPosInfoDto() 从对象池中获取PosInfoDto
func GetPosInfoDto() *PosInfoDto {
	return poolPosInfoDto.Get().(*PosInfoDto)
}

// ReleasePosInfoDto 释放PosInfoDto
func ReleasePosInfoDto(v *PosInfoDto) {
	v.Support4Decimal = ""
	v.CounterNo = ""
	v.StoreNo = ""
	poolPosInfoDto.Put(v)
}
