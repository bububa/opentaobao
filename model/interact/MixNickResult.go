package interact

import (
	"sync"
)

// MixNickResult 结构体
type MixNickResult struct {
	// 互动混淆nick
	PlayMixnick string `json:"play_mixnick,omitempty" xml:"play_mixnick,omitempty"`
}

var poolMixNickResult = sync.Pool{
	New: func() any {
		return new(MixNickResult)
	},
}

// GetMixNickResult() 从对象池中获取MixNickResult
func GetMixNickResult() *MixNickResult {
	return poolMixNickResult.Get().(*MixNickResult)
}

// ReleaseMixNickResult 释放MixNickResult
func ReleaseMixNickResult(v *MixNickResult) {
	v.PlayMixnick = ""
	poolMixNickResult.Put(v)
}
