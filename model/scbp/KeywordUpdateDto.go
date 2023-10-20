package scbp

import (
	"sync"
)

// KeywordUpdateDto 结构体
type KeywordUpdateDto struct {
	// 要改的价格，单位元
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 词id
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
}

var poolKeywordUpdateDto = sync.Pool{
	New: func() any {
		return new(KeywordUpdateDto)
	},
}

// GetKeywordUpdateDto() 从对象池中获取KeywordUpdateDto
func GetKeywordUpdateDto() *KeywordUpdateDto {
	return poolKeywordUpdateDto.Get().(*KeywordUpdateDto)
}

// ReleaseKeywordUpdateDto 释放KeywordUpdateDto
func ReleaseKeywordUpdateDto(v *KeywordUpdateDto) {
	v.Value = ""
	v.KeywordId = 0
	poolKeywordUpdateDto.Put(v)
}
