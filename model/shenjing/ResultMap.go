package shenjing

import (
	"sync"
)

// ResultMap 结构体
type ResultMap struct {
	// 图片URL
	PhotoUrl string `json:"photo_url,omitempty" xml:"photo_url,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
}

var poolResultMap = sync.Pool{
	New: func() any {
		return new(ResultMap)
	},
}

// GetResultMap() 从对象池中获取ResultMap
func GetResultMap() *ResultMap {
	return poolResultMap.Get().(*ResultMap)
}

// ReleaseResultMap 释放ResultMap
func ReleaseResultMap(v *ResultMap) {
	v.PhotoUrl = ""
	v.CampusName = ""
	poolResultMap.Put(v)
}
