package singletreasure

import (
	"sync"
)

// TaobaoSingletreasureActivityItemBatchaddResultDto 结构体
type TaobaoSingletreasureActivityItemBatchaddResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *ItemProcessErrorResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSingletreasureActivityItemBatchaddResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemBatchaddResultDto)
	},
}

// GetTaobaoSingletreasureActivityItemBatchaddResultDto() 从对象池中获取TaobaoSingletreasureActivityItemBatchaddResultDto
func GetTaobaoSingletreasureActivityItemBatchaddResultDto() *TaobaoSingletreasureActivityItemBatchaddResultDto {
	return poolTaobaoSingletreasureActivityItemBatchaddResultDto.Get().(*TaobaoSingletreasureActivityItemBatchaddResultDto)
}

// ReleaseTaobaoSingletreasureActivityItemBatchaddResultDto 释放TaobaoSingletreasureActivityItemBatchaddResultDto
func ReleaseTaobaoSingletreasureActivityItemBatchaddResultDto(v *TaobaoSingletreasureActivityItemBatchaddResultDto) {
	v.Message = ""
	v.Data = nil
	v.Code = 0
	v.Success = false
	poolTaobaoSingletreasureActivityItemBatchaddResultDto.Put(v)
}
