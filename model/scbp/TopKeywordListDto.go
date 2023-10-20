package scbp

import (
	"sync"
)

// TopKeywordListDto 结构体
type TopKeywordListDto struct {
	// 关键词列表
	KeywordList []string `json:"keyword_list,omitempty" xml:"keyword_list>string,omitempty"`
}

var poolTopKeywordListDto = sync.Pool{
	New: func() any {
		return new(TopKeywordListDto)
	},
}

// GetTopKeywordListDto() 从对象池中获取TopKeywordListDto
func GetTopKeywordListDto() *TopKeywordListDto {
	return poolTopKeywordListDto.Get().(*TopKeywordListDto)
}

// ReleaseTopKeywordListDto 释放TopKeywordListDto
func ReleaseTopKeywordListDto(v *TopKeywordListDto) {
	v.KeywordList = v.KeywordList[:0]
	poolTopKeywordListDto.Put(v)
}
