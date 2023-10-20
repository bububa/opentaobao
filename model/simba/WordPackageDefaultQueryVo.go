package simba

import (
	"sync"
)

// WordPackageDefaultQueryVo 结构体
type WordPackageDefaultQueryVo struct {
	// 宝贝id集合
	MaterialIdList []int64 `json:"material_id_list,omitempty" xml:"material_id_list>int64,omitempty"`
}

var poolWordPackageDefaultQueryVo = sync.Pool{
	New: func() any {
		return new(WordPackageDefaultQueryVo)
	},
}

// GetWordPackageDefaultQueryVo() 从对象池中获取WordPackageDefaultQueryVo
func GetWordPackageDefaultQueryVo() *WordPackageDefaultQueryVo {
	return poolWordPackageDefaultQueryVo.Get().(*WordPackageDefaultQueryVo)
}

// ReleaseWordPackageDefaultQueryVo 释放WordPackageDefaultQueryVo
func ReleaseWordPackageDefaultQueryVo(v *WordPackageDefaultQueryVo) {
	v.MaterialIdList = v.MaterialIdList[:0]
	poolWordPackageDefaultQueryVo.Put(v)
}
