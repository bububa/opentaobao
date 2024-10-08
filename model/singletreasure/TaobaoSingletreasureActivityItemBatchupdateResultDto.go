package singletreasure

import (
	"sync"
)

// TaobaoSingletreasureActivityItemBatchupdateResultDto 结构体
type TaobaoSingletreasureActivityItemBatchupdateResultDto struct {
	// 返回的描述信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回所有的处理错误的信息
	Data *ItemProcessErrorResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 处理结果
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSingletreasureActivityItemBatchupdateResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemBatchupdateResultDto)
	},
}

// GetTaobaoSingletreasureActivityItemBatchupdateResultDto() 从对象池中获取TaobaoSingletreasureActivityItemBatchupdateResultDto
func GetTaobaoSingletreasureActivityItemBatchupdateResultDto() *TaobaoSingletreasureActivityItemBatchupdateResultDto {
	return poolTaobaoSingletreasureActivityItemBatchupdateResultDto.Get().(*TaobaoSingletreasureActivityItemBatchupdateResultDto)
}

// ReleaseTaobaoSingletreasureActivityItemBatchupdateResultDto 释放TaobaoSingletreasureActivityItemBatchupdateResultDto
func ReleaseTaobaoSingletreasureActivityItemBatchupdateResultDto(v *TaobaoSingletreasureActivityItemBatchupdateResultDto) {
	v.Message = ""
	v.Data = nil
	v.Code = 0
	v.Success = false
	poolTaobaoSingletreasureActivityItemBatchupdateResultDto.Put(v)
}
