package category

import (
	"sync"
)

// Brand 结构体
type Brand struct {
	// vid的值
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 品牌的属性名
	PropName string `json:"prop_name,omitempty" xml:"prop_name,omitempty"`
	// 对应属性的 pid:vid 串中的vid
	Vid int64 `json:"vid,omitempty" xml:"vid,omitempty"`
	// 品牌的属性id
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
}

var poolBrand = sync.Pool{
	New: func() any {
		return new(Brand)
	},
}

// GetBrand() 从对象池中获取Brand
func GetBrand() *Brand {
	return poolBrand.Get().(*Brand)
}

// ReleaseBrand 释放Brand
func ReleaseBrand(v *Brand) {
	v.Name = ""
	v.PropName = ""
	v.Vid = 0
	v.Pid = 0
	poolBrand.Put(v)
}
