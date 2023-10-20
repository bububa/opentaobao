package aesolution

import (
	"sync"
)

// GlobalSubject 结构体
type GlobalSubject struct {
	// locale of the subject
	Locale string `json:"locale,omitempty" xml:"locale,omitempty"`
	// subject
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
}

var poolGlobalSubject = sync.Pool{
	New: func() any {
		return new(GlobalSubject)
	},
}

// GetGlobalSubject() 从对象池中获取GlobalSubject
func GetGlobalSubject() *GlobalSubject {
	return poolGlobalSubject.Get().(*GlobalSubject)
}

// ReleaseGlobalSubject 释放GlobalSubject
func ReleaseGlobalSubject(v *GlobalSubject) {
	v.Locale = ""
	v.Subject = ""
	poolGlobalSubject.Put(v)
}
