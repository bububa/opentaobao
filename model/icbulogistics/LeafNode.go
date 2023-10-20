package icbulogistics

import (
	"sync"
)

// LeafNode 结构体
type LeafNode struct {
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolLeafNode = sync.Pool{
	New: func() any {
		return new(LeafNode)
	},
}

// GetLeafNode() 从对象池中获取LeafNode
func GetLeafNode() *LeafNode {
	return poolLeafNode.Get().(*LeafNode)
}

// ReleaseLeafNode 释放LeafNode
func ReleaseLeafNode(v *LeafNode) {
	v.Code = ""
	v.Name = ""
	poolLeafNode.Put(v)
}
