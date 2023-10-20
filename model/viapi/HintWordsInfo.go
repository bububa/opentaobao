package viapi

import (
	"sync"
)

// HintWordsInfo 结构体
type HintWordsInfo struct {
	// 图片中文字命中的风险关键词内容
	Context string `json:"context,omitempty" xml:"context,omitempty"`
}

var poolHintWordsInfo = sync.Pool{
	New: func() any {
		return new(HintWordsInfo)
	},
}

// GetHintWordsInfo() 从对象池中获取HintWordsInfo
func GetHintWordsInfo() *HintWordsInfo {
	return poolHintWordsInfo.Get().(*HintWordsInfo)
}

// ReleaseHintWordsInfo 释放HintWordsInfo
func ReleaseHintWordsInfo(v *HintWordsInfo) {
	v.Context = ""
	poolHintWordsInfo.Put(v)
}
