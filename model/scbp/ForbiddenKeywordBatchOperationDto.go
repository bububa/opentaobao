package scbp

import (
	"sync"
)

// ForbiddenKeywordBatchOperationDto 结构体
type ForbiddenKeywordBatchOperationDto struct {
	// 请求参数
	ForbiddenKeywordOperationList []ForbiddenKeywordOperation `json:"forbidden_keyword_operation_list,omitempty" xml:"forbidden_keyword_operation_list>forbidden_keyword_operation,omitempty"`
}

var poolForbiddenKeywordBatchOperationDto = sync.Pool{
	New: func() any {
		return new(ForbiddenKeywordBatchOperationDto)
	},
}

// GetForbiddenKeywordBatchOperationDto() 从对象池中获取ForbiddenKeywordBatchOperationDto
func GetForbiddenKeywordBatchOperationDto() *ForbiddenKeywordBatchOperationDto {
	return poolForbiddenKeywordBatchOperationDto.Get().(*ForbiddenKeywordBatchOperationDto)
}

// ReleaseForbiddenKeywordBatchOperationDto 释放ForbiddenKeywordBatchOperationDto
func ReleaseForbiddenKeywordBatchOperationDto(v *ForbiddenKeywordBatchOperationDto) {
	v.ForbiddenKeywordOperationList = v.ForbiddenKeywordOperationList[:0]
	poolForbiddenKeywordBatchOperationDto.Put(v)
}
