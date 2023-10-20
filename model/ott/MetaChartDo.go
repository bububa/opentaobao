package ott

import (
	"sync"
)

// MetaChartDo 结构体
type MetaChartDo struct {
	// 排行标题
	Titles []string `json:"titles,omitempty" xml:"titles>string,omitempty"`
	// 排行类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolMetaChartDo = sync.Pool{
	New: func() any {
		return new(MetaChartDo)
	},
}

// GetMetaChartDo() 从对象池中获取MetaChartDo
func GetMetaChartDo() *MetaChartDo {
	return poolMetaChartDo.Get().(*MetaChartDo)
}

// ReleaseMetaChartDo 释放MetaChartDo
func ReleaseMetaChartDo(v *MetaChartDo) {
	v.Titles = v.Titles[:0]
	v.Type = ""
	poolMetaChartDo.Put(v)
}
