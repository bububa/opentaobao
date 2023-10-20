package aesolution

import (
	"sync"
)

// GlobalDescription 结构体
type GlobalDescription struct {
	// locale of the descripiton
	Locale string `json:"locale,omitempty" xml:"locale,omitempty"`
	// mobile detail
	MobileDetail string `json:"mobile_detail,omitempty" xml:"mobile_detail,omitempty"`
	// web detail
	WebDetail string `json:"web_detail,omitempty" xml:"web_detail,omitempty"`
}

var poolGlobalDescription = sync.Pool{
	New: func() any {
		return new(GlobalDescription)
	},
}

// GetGlobalDescription() 从对象池中获取GlobalDescription
func GetGlobalDescription() *GlobalDescription {
	return poolGlobalDescription.Get().(*GlobalDescription)
}

// ReleaseGlobalDescription 释放GlobalDescription
func ReleaseGlobalDescription(v *GlobalDescription) {
	v.Locale = ""
	v.MobileDetail = ""
	v.WebDetail = ""
	poolGlobalDescription.Put(v)
}
