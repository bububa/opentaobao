package campus

import (
	"sync"
)

// ControllerQuery 结构体
type ControllerQuery struct {
	// 序列号
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
}

var poolControllerQuery = sync.Pool{
	New: func() any {
		return new(ControllerQuery)
	},
}

// GetControllerQuery() 从对象池中获取ControllerQuery
func GetControllerQuery() *ControllerQuery {
	return poolControllerQuery.Get().(*ControllerQuery)
}

// ReleaseControllerQuery 释放ControllerQuery
func ReleaseControllerQuery(v *ControllerQuery) {
	v.SnNo = ""
	poolControllerQuery.Put(v)
}
