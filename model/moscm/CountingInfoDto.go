package moscm

import (
	"sync"
)

// CountingInfoDto 结构体
type CountingInfoDto struct {
	// 专柜id
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 外部专柜id(两个专柜号至少填写一个。如果同时填写，以counterId为准)
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 备注信息（必填）
	Remarks string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

var poolCountingInfoDto = sync.Pool{
	New: func() any {
		return new(CountingInfoDto)
	},
}

// GetCountingInfoDto() 从对象池中获取CountingInfoDto
func GetCountingInfoDto() *CountingInfoDto {
	return poolCountingInfoDto.Get().(*CountingInfoDto)
}

// ReleaseCountingInfoDto 释放CountingInfoDto
func ReleaseCountingInfoDto(v *CountingInfoDto) {
	v.CounterId = ""
	v.OutCounterId = ""
	v.Remarks = ""
	poolCountingInfoDto.Put(v)
}
