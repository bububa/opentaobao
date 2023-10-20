package simba

import (
	"sync"
)

// SuggestWordPackageVo 结构体
type SuggestWordPackageVo struct {
	// 主题词
	ThemeWordList []string `json:"theme_word_list,omitempty" xml:"theme_word_list>string,omitempty"`
	// 关键词举例
	SimpleWordList []string `json:"simple_word_list,omitempty" xml:"simple_word_list>string,omitempty"`
	// 词包名称
	WordPackageName string `json:"word_package_name,omitempty" xml:"word_package_name,omitempty"`
	// 建议出价
	BidPrice string `json:"bid_price,omitempty" xml:"bid_price,omitempty"`
	// 预估展现
	Impression string `json:"impression,omitempty" xml:"impression,omitempty"`
	// 流量放大系数
	MultiFactor string `json:"multi_factor,omitempty" xml:"multi_factor,omitempty"`
	// 词包id
	WordPackageId int64 `json:"word_package_id,omitempty" xml:"word_package_id,omitempty"`
	// 词包类型,0:流量智选,20:卖点词包
	WordPackageType int64 `json:"word_package_type,omitempty" xml:"word_package_type,omitempty"`
	// 相关性,1:差,2:中,3:好
	RelevanceType int64 `json:"relevance_type,omitempty" xml:"relevance_type,omitempty"`
}

var poolSuggestWordPackageVo = sync.Pool{
	New: func() any {
		return new(SuggestWordPackageVo)
	},
}

// GetSuggestWordPackageVo() 从对象池中获取SuggestWordPackageVo
func GetSuggestWordPackageVo() *SuggestWordPackageVo {
	return poolSuggestWordPackageVo.Get().(*SuggestWordPackageVo)
}

// ReleaseSuggestWordPackageVo 释放SuggestWordPackageVo
func ReleaseSuggestWordPackageVo(v *SuggestWordPackageVo) {
	v.ThemeWordList = v.ThemeWordList[:0]
	v.SimpleWordList = v.SimpleWordList[:0]
	v.WordPackageName = ""
	v.BidPrice = ""
	v.Impression = ""
	v.MultiFactor = ""
	v.WordPackageId = 0
	v.WordPackageType = 0
	v.RelevanceType = 0
	poolSuggestWordPackageVo.Put(v)
}
