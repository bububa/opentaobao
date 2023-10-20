package singletreasure

import (
	"sync"
)

// TaobaoSingletreasureActivityNameQueryResultDto 结构体
type TaobaoSingletreasureActivityNameQueryResultDto struct {
	// data结果
	DataList []ActivityNameCategoryDto `json:"data_list,omitempty" xml:"data_list>activity_name_category_dto,omitempty"`
	// 请求返回描述信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoSingletreasureActivityNameQueryResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSingletreasureActivityNameQueryResultDto)
	},
}

// GetTaobaoSingletreasureActivityNameQueryResultDto() 从对象池中获取TaobaoSingletreasureActivityNameQueryResultDto
func GetTaobaoSingletreasureActivityNameQueryResultDto() *TaobaoSingletreasureActivityNameQueryResultDto {
	return poolTaobaoSingletreasureActivityNameQueryResultDto.Get().(*TaobaoSingletreasureActivityNameQueryResultDto)
}

// ReleaseTaobaoSingletreasureActivityNameQueryResultDto 释放TaobaoSingletreasureActivityNameQueryResultDto
func ReleaseTaobaoSingletreasureActivityNameQueryResultDto(v *TaobaoSingletreasureActivityNameQueryResultDto) {
	v.DataList = v.DataList[:0]
	v.Message = ""
	v.Code = 0
	v.Success = false
	poolTaobaoSingletreasureActivityNameQueryResultDto.Put(v)
}
