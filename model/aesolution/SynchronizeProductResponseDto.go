package aesolution

import (
	"sync"
)

// SynchronizeProductResponseDto 结构体
type SynchronizeProductResponseDto struct {
	// error code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolSynchronizeProductResponseDto = sync.Pool{
	New: func() any {
		return new(SynchronizeProductResponseDto)
	},
}

// GetSynchronizeProductResponseDto() 从对象池中获取SynchronizeProductResponseDto
func GetSynchronizeProductResponseDto() *SynchronizeProductResponseDto {
	return poolSynchronizeProductResponseDto.Get().(*SynchronizeProductResponseDto)
}

// ReleaseSynchronizeProductResponseDto 释放SynchronizeProductResponseDto
func ReleaseSynchronizeProductResponseDto(v *SynchronizeProductResponseDto) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.ProductId = 0
	poolSynchronizeProductResponseDto.Put(v)
}
