package yunosappstore

import (
	"sync"
)

// Callbacks 结构体
type Callbacks struct {
	// 点击
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 曝光
	View string `json:"view,omitempty" xml:"view,omitempty"`
	// 竞价成功
	Deal string `json:"deal,omitempty" xml:"deal,omitempty"`
}

var poolCallbacks = sync.Pool{
	New: func() any {
		return new(Callbacks)
	},
}

// GetCallbacks() 从对象池中获取Callbacks
func GetCallbacks() *Callbacks {
	return poolCallbacks.Get().(*Callbacks)
}

// ReleaseCallbacks 释放Callbacks
func ReleaseCallbacks(v *Callbacks) {
	v.Click = ""
	v.View = ""
	v.Deal = ""
	poolCallbacks.Put(v)
}
