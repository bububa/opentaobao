package btrip

import (
	"sync"
)

// OpenSuggestRq 结构体
type OpenSuggestRq struct {
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 搜索关键字
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 0国内机场，2国内机场+临近机场，3国际机场
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenSuggestRq = sync.Pool{
	New: func() any {
		return new(OpenSuggestRq)
	},
}

// GetOpenSuggestRq() 从对象池中获取OpenSuggestRq
func GetOpenSuggestRq() *OpenSuggestRq {
	return poolOpenSuggestRq.Get().(*OpenSuggestRq)
}

// ReleaseOpenSuggestRq 释放OpenSuggestRq
func ReleaseOpenSuggestRq(v *OpenSuggestRq) {
	v.CorpId = ""
	v.Keyword = ""
	v.Type = 0
	v.Version = 0
	poolOpenSuggestRq.Put(v)
}
