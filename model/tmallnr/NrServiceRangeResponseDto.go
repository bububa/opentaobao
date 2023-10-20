package tmallnr

import (
	"sync"
)

// NrServiceRangeResponseDto 结构体
type NrServiceRangeResponseDto struct {
	// 围栏数据
	Points []Points `json:"points,omitempty" xml:"points>points,omitempty"`
}

var poolNrServiceRangeResponseDto = sync.Pool{
	New: func() any {
		return new(NrServiceRangeResponseDto)
	},
}

// GetNrServiceRangeResponseDto() 从对象池中获取NrServiceRangeResponseDto
func GetNrServiceRangeResponseDto() *NrServiceRangeResponseDto {
	return poolNrServiceRangeResponseDto.Get().(*NrServiceRangeResponseDto)
}

// ReleaseNrServiceRangeResponseDto 释放NrServiceRangeResponseDto
func ReleaseNrServiceRangeResponseDto(v *NrServiceRangeResponseDto) {
	v.Points = v.Points[:0]
	poolNrServiceRangeResponseDto.Put(v)
}
