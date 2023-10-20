package simba

import (
	"sync"
)

// SubCrowdVo 结构体
type SubCrowdVo struct {
	// 子定向名称
	SubcrowdName string `json:"subcrowd_name,omitempty" xml:"subcrowd_name,omitempty"`
	// 子定向值
	SubcrowdValue string `json:"subcrowd_value,omitempty" xml:"subcrowd_value,omitempty"`
}

var poolSubCrowdVo = sync.Pool{
	New: func() any {
		return new(SubCrowdVo)
	},
}

// GetSubCrowdVo() 从对象池中获取SubCrowdVo
func GetSubCrowdVo() *SubCrowdVo {
	return poolSubCrowdVo.Get().(*SubCrowdVo)
}

// ReleaseSubCrowdVo 释放SubCrowdVo
func ReleaseSubCrowdVo(v *SubCrowdVo) {
	v.SubcrowdName = ""
	v.SubcrowdValue = ""
	poolSubCrowdVo.Put(v)
}
