package singletreasure

import (
	"sync"
)

// TaobaoSingletreasureActivityCreateResultDto 结构体
type TaobaoSingletreasureActivityCreateResultDto struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 新建套餐 id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 错误编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 系统执行成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSingletreasureActivityCreateResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityCreateResultDto)
	},
}

// GetTaobaoSingletreasureActivityCreateResultDto() 从对象池中获取TaobaoSingletreasureActivityCreateResultDto
func GetTaobaoSingletreasureActivityCreateResultDto() *TaobaoSingletreasureActivityCreateResultDto {
	return poolTaobaoSingletreasureActivityCreateResultDto.Get().(*TaobaoSingletreasureActivityCreateResultDto)
}

// ReleaseTaobaoSingletreasureActivityCreateResultDto 释放TaobaoSingletreasureActivityCreateResultDto
func ReleaseTaobaoSingletreasureActivityCreateResultDto(v *TaobaoSingletreasureActivityCreateResultDto) {
	v.Message = ""
	v.Data = 0
	v.Code = 0
	v.Success = false
	poolTaobaoSingletreasureActivityCreateResultDto.Put(v)
}
