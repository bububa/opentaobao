package viapi

import (
	"sync"
)

// Frame 结构体
type Frame struct {
	// 被截断的图片的临时访问url，地址有效期是5分钟
	URL string `json:"u_r_l,omitempty" xml:"u_r_l,omitempty"`
	// 置信度，仅作参考，不建议使用
	Rate int64 `json:"rate,omitempty" xml:"rate,omitempty"`
}

var poolFrame = sync.Pool{
	New: func() any {
		return new(Frame)
	},
}

// GetFrame() 从对象池中获取Frame
func GetFrame() *Frame {
	return poolFrame.Get().(*Frame)
}

// ReleaseFrame 释放Frame
func ReleaseFrame(v *Frame) {
	v.URL = ""
	v.Rate = 0
	poolFrame.Put(v)
}
