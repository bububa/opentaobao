package idleparttime

import (
	"sync"
)

// PartTimeJobTransNodes 结构体
type PartTimeJobTransNodes struct {
	// 节点描述
	NodeDescription string `json:"node_description,omitempty" xml:"node_description,omitempty"`
	// 节点创建时间
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

var poolPartTimeJobTransNodes = sync.Pool{
	New: func() any {
		return new(PartTimeJobTransNodes)
	},
}

// GetPartTimeJobTransNodes() 从对象池中获取PartTimeJobTransNodes
func GetPartTimeJobTransNodes() *PartTimeJobTransNodes {
	return poolPartTimeJobTransNodes.Get().(*PartTimeJobTransNodes)
}

// ReleasePartTimeJobTransNodes 释放PartTimeJobTransNodes
func ReleasePartTimeJobTransNodes(v *PartTimeJobTransNodes) {
	v.NodeDescription = ""
	v.CreateTime = 0
	poolPartTimeJobTransNodes.Put(v)
}
