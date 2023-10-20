package singletreasure

import (
	"sync"
)

// TaobaoSingletreasureActivityItemUpdateResultDto 结构体
type TaobaoSingletreasureActivityItemUpdateResultDto struct {
	// 请求返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回所有的处理错误的信息
	Data *ItemProcessErrorResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 请求返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSingletreasureActivityItemUpdateResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemUpdateResultDto)
	},
}

// GetTaobaoSingletreasureActivityItemUpdateResultDto() 从对象池中获取TaobaoSingletreasureActivityItemUpdateResultDto
func GetTaobaoSingletreasureActivityItemUpdateResultDto() *TaobaoSingletreasureActivityItemUpdateResultDto {
	return poolTaobaoSingletreasureActivityItemUpdateResultDto.Get().(*TaobaoSingletreasureActivityItemUpdateResultDto)
}

// ReleaseTaobaoSingletreasureActivityItemUpdateResultDto 释放TaobaoSingletreasureActivityItemUpdateResultDto
func ReleaseTaobaoSingletreasureActivityItemUpdateResultDto(v *TaobaoSingletreasureActivityItemUpdateResultDto) {
	v.Message = ""
	v.Data = nil
	v.Code = 0
	v.Success = false
	poolTaobaoSingletreasureActivityItemUpdateResultDto.Put(v)
}
