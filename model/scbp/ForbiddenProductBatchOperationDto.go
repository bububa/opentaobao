package scbp

import (
	"sync"
)

// ForbiddenProductBatchOperationDto 结构体
type ForbiddenProductBatchOperationDto struct {
	// 查询条件
	ForbiddenProductOperationList []ForbiddenProductOperationDto `json:"forbidden_product_operation_list,omitempty" xml:"forbidden_product_operation_list>forbidden_product_operation_dto,omitempty"`
}

var poolForbiddenProductBatchOperationDto = sync.Pool{
	New: func() any {
		return new(ForbiddenProductBatchOperationDto)
	},
}

// GetForbiddenProductBatchOperationDto() 从对象池中获取ForbiddenProductBatchOperationDto
func GetForbiddenProductBatchOperationDto() *ForbiddenProductBatchOperationDto {
	return poolForbiddenProductBatchOperationDto.Get().(*ForbiddenProductBatchOperationDto)
}

// ReleaseForbiddenProductBatchOperationDto 释放ForbiddenProductBatchOperationDto
func ReleaseForbiddenProductBatchOperationDto(v *ForbiddenProductBatchOperationDto) {
	v.ForbiddenProductOperationList = v.ForbiddenProductOperationList[:0]
	poolForbiddenProductBatchOperationDto.Put(v)
}
