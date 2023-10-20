package scbp

import (
	"sync"
)

// ForbiddenProductOperationDto 结构体
type ForbiddenProductOperationDto struct {
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolForbiddenProductOperationDto = sync.Pool{
	New: func() any {
		return new(ForbiddenProductOperationDto)
	},
}

// GetForbiddenProductOperationDto() 从对象池中获取ForbiddenProductOperationDto
func GetForbiddenProductOperationDto() *ForbiddenProductOperationDto {
	return poolForbiddenProductOperationDto.Get().(*ForbiddenProductOperationDto)
}

// ReleaseForbiddenProductOperationDto 释放ForbiddenProductOperationDto
func ReleaseForbiddenProductOperationDto(v *ForbiddenProductOperationDto) {
	v.ProductId = 0
	poolForbiddenProductOperationDto.Put(v)
}
