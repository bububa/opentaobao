package security

import (
	"sync"
)

// RpGrade 结构体
type RpGrade struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// level
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}

var poolRpGrade = sync.Pool{
	New: func() any {
		return new(RpGrade)
	},
}

// GetRpGrade() 从对象池中获取RpGrade
func GetRpGrade() *RpGrade {
	return poolRpGrade.Get().(*RpGrade)
}

// ReleaseRpGrade 释放RpGrade
func ReleaseRpGrade(v *RpGrade) {
	v.Desc = ""
	v.Name = ""
	v.Level = 0
	poolRpGrade.Put(v)
}
