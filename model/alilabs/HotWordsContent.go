package alilabs

import (
	"sync"
)

// HotWordsContent 结构体
type HotWordsContent struct {
	// 热词列表
	Words []string `json:"words,omitempty" xml:"words>string,omitempty"`
}

var poolHotWordsContent = sync.Pool{
	New: func() any {
		return new(HotWordsContent)
	},
}

// GetHotWordsContent() 从对象池中获取HotWordsContent
func GetHotWordsContent() *HotWordsContent {
	return poolHotWordsContent.Get().(*HotWordsContent)
}

// ReleaseHotWordsContent 释放HotWordsContent
func ReleaseHotWordsContent(v *HotWordsContent) {
	v.Words = v.Words[:0]
	poolHotWordsContent.Put(v)
}
