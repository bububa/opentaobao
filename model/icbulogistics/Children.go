package icbulogistics

import (
	"sync"
)

// Children 结构体
type Children struct {
	// 商品特性列表对象
	Children []LeafNode `json:"children,omitempty" xml:"children>leaf_node,omitempty"`
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolChildren = sync.Pool{
	New: func() any {
		return new(Children)
	},
}

// GetChildren() 从对象池中获取Children
func GetChildren() *Children {
	return poolChildren.Get().(*Children)
}

// ReleaseChildren 释放Children
func ReleaseChildren(v *Children) {
	v.Children = v.Children[:0]
	v.Code = ""
	v.Name = ""
	poolChildren.Put(v)
}
