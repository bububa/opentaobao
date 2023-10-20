package scbp

import (
	"sync"
)

// KeywordErrorResultDto 结构体
type KeywordErrorResultDto struct {
	// reason
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// keywordId
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
}

var poolKeywordErrorResultDto = sync.Pool{
	New: func() any {
		return new(KeywordErrorResultDto)
	},
}

// GetKeywordErrorResultDto() 从对象池中获取KeywordErrorResultDto
func GetKeywordErrorResultDto() *KeywordErrorResultDto {
	return poolKeywordErrorResultDto.Get().(*KeywordErrorResultDto)
}

// ReleaseKeywordErrorResultDto 释放KeywordErrorResultDto
func ReleaseKeywordErrorResultDto(v *KeywordErrorResultDto) {
	v.Reason = ""
	v.Value = ""
	v.KeywordId = 0
	poolKeywordErrorResultDto.Put(v)
}
