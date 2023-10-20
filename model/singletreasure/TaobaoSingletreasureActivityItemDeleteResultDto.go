package singletreasure

import (
	"sync"
)

// TaobaoSingletreasureActivityItemDeleteResultDto 结构体
type TaobaoSingletreasureActivityItemDeleteResultDto struct {
	// 请求返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回所有的处理错误的信息
	Data *ItemProcessErrorResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 请求返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSingletreasureActivityItemDeleteResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityItemDeleteResultDto)
	},
}

// GetTaobaoSingletreasureActivityItemDeleteResultDto() 从对象池中获取TaobaoSingletreasureActivityItemDeleteResultDto
func GetTaobaoSingletreasureActivityItemDeleteResultDto() *TaobaoSingletreasureActivityItemDeleteResultDto {
	return poolTaobaoSingletreasureActivityItemDeleteResultDto.Get().(*TaobaoSingletreasureActivityItemDeleteResultDto)
}

// ReleaseTaobaoSingletreasureActivityItemDeleteResultDto 释放TaobaoSingletreasureActivityItemDeleteResultDto
func ReleaseTaobaoSingletreasureActivityItemDeleteResultDto(v *TaobaoSingletreasureActivityItemDeleteResultDto) {
	v.Message = ""
	v.Data = nil
	v.Code = 0
	v.Success = false
	poolTaobaoSingletreasureActivityItemDeleteResultDto.Put(v)
}
