package scbp

import (
	"sync"
)

// ForbiddenKeywordOperationDto 结构体
type ForbiddenKeywordOperationDto struct {
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

var poolForbiddenKeywordOperationDto = sync.Pool{
	New: func() any {
		return new(ForbiddenKeywordOperationDto)
	},
}

// GetForbiddenKeywordOperationDto() 从对象池中获取ForbiddenKeywordOperationDto
func GetForbiddenKeywordOperationDto() *ForbiddenKeywordOperationDto {
	return poolForbiddenKeywordOperationDto.Get().(*ForbiddenKeywordOperationDto)
}

// ReleaseForbiddenKeywordOperationDto 释放ForbiddenKeywordOperationDto
func ReleaseForbiddenKeywordOperationDto(v *ForbiddenKeywordOperationDto) {
	v.Keyword = ""
	poolForbiddenKeywordOperationDto.Put(v)
}
