package simba

import (
	"sync"
)

// WordPackageSuggestItemVo 结构体
type WordPackageSuggestItemVo struct {
	// 推荐的词包列表
	WordPackageList []SuggestWordPackageVo `json:"word_package_list,omitempty" xml:"word_package_list>suggest_word_package_vo,omitempty"`
	// 宝贝id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
}

var poolWordPackageSuggestItemVo = sync.Pool{
	New: func() any {
		return new(WordPackageSuggestItemVo)
	},
}

// GetWordPackageSuggestItemVo() 从对象池中获取WordPackageSuggestItemVo
func GetWordPackageSuggestItemVo() *WordPackageSuggestItemVo {
	return poolWordPackageSuggestItemVo.Get().(*WordPackageSuggestItemVo)
}

// ReleaseWordPackageSuggestItemVo 释放WordPackageSuggestItemVo
func ReleaseWordPackageSuggestItemVo(v *WordPackageSuggestItemVo) {
	v.WordPackageList = v.WordPackageList[:0]
	v.MaterialId = 0
	poolWordPackageSuggestItemVo.Put(v)
}
