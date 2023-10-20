package scbp

import (
	"sync"
)

// KeywordInfo 结构体
type KeywordInfo struct {
	// 关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 关键词id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
}

var poolKeywordInfo = sync.Pool{
	New: func() any {
		return new(KeywordInfo)
	},
}

// GetKeywordInfo() 从对象池中获取KeywordInfo
func GetKeywordInfo() *KeywordInfo {
	return poolKeywordInfo.Get().(*KeywordInfo)
}

// ReleaseKeywordInfo 释放KeywordInfo
func ReleaseKeywordInfo(v *KeywordInfo) {
	v.Word = ""
	v.Price = ""
	v.Id = 0
	v.OnlineStatus = 0
	poolKeywordInfo.Put(v)
}
