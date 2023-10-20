package mydata

import (
	"sync"
)

// KeywordEffectEntity 结构体
type KeywordEffectEntity struct {
	// 词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

var poolKeywordEffectEntity = sync.Pool{
	New: func() any {
		return new(KeywordEffectEntity)
	},
}

// GetKeywordEffectEntity() 从对象池中获取KeywordEffectEntity
func GetKeywordEffectEntity() *KeywordEffectEntity {
	return poolKeywordEffectEntity.Get().(*KeywordEffectEntity)
}

// ReleaseKeywordEffectEntity 释放KeywordEffectEntity
func ReleaseKeywordEffectEntity(v *KeywordEffectEntity) {
	v.Keyword = ""
	poolKeywordEffectEntity.Put(v)
}
