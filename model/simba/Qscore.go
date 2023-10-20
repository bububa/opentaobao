package simba

import (
	"sync"
)

// Qscore 结构体
type Qscore struct {
	// 词质量得分列表
	KeywordQscoreList []KeywordQscore `json:"keyword_qscore_list,omitempty" xml:"keyword_qscore_list>keyword_qscore,omitempty"`
	// 类目出价质量得分
	CatmatchQscore string `json:"catmatch_qscore,omitempty" xml:"catmatch_qscore,omitempty"`
}

var poolQscore = sync.Pool{
	New: func() any {
		return new(Qscore)
	},
}

// GetQscore() 从对象池中获取Qscore
func GetQscore() *Qscore {
	return poolQscore.Get().(*Qscore)
}

// ReleaseQscore 释放Qscore
func ReleaseQscore(v *Qscore) {
	v.KeywordQscoreList = v.KeywordQscoreList[:0]
	v.CatmatchQscore = ""
	poolQscore.Put(v)
}
