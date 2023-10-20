package travel

import (
	"sync"
)

// TopElementResult 结构体
type TopElementResult struct {
	// 元素的外部商家编码
	ElementOuterId string `json:"element_outer_id,omitempty" xml:"element_outer_id,omitempty"`
	// 创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 删除时间
	Deleted string `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolTopElementResult = sync.Pool{
	New: func() any {
		return new(TopElementResult)
	},
}

// GetTopElementResult() 从对象池中获取TopElementResult
func GetTopElementResult() *TopElementResult {
	return poolTopElementResult.Get().(*TopElementResult)
}

// ReleaseTopElementResult 释放TopElementResult
func ReleaseTopElementResult(v *TopElementResult) {
	v.ElementOuterId = ""
	v.Created = ""
	v.Modified = ""
	v.Deleted = ""
	poolTopElementResult.Put(v)
}
