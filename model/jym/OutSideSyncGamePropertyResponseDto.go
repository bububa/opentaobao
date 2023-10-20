package jym

import (
	"sync"
)

// OutSideSyncGamePropertyResponseDto 结构体
type OutSideSyncGamePropertyResponseDto struct {
	// true - 成功，false - 失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolOutSideSyncGamePropertyResponseDto = sync.Pool{
	New: func() any {
		return new(OutSideSyncGamePropertyResponseDto)
	},
}

// GetOutSideSyncGamePropertyResponseDto() 从对象池中获取OutSideSyncGamePropertyResponseDto
func GetOutSideSyncGamePropertyResponseDto() *OutSideSyncGamePropertyResponseDto {
	return poolOutSideSyncGamePropertyResponseDto.Get().(*OutSideSyncGamePropertyResponseDto)
}

// ReleaseOutSideSyncGamePropertyResponseDto 释放OutSideSyncGamePropertyResponseDto
func ReleaseOutSideSyncGamePropertyResponseDto(v *OutSideSyncGamePropertyResponseDto) {
	v.Result = false
	poolOutSideSyncGamePropertyResponseDto.Put(v)
}
