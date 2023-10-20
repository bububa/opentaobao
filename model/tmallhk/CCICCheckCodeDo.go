package tmallhk

import (
	"sync"
)

// CCICCheckCodeDo 结构体
type CCICCheckCodeDo struct {
	// enterUrl
	EnterUrl string `json:"enter_url,omitempty" xml:"enter_url,omitempty"`
	// enterable
	Enterable bool `json:"enterable,omitempty" xml:"enterable,omitempty"`
}

var poolCCICCheckCodeDo = sync.Pool{
	New: func() any {
		return new(CCICCheckCodeDo)
	},
}

// GetCCICCheckCodeDo() 从对象池中获取CCICCheckCodeDo
func GetCCICCheckCodeDo() *CCICCheckCodeDo {
	return poolCCICCheckCodeDo.Get().(*CCICCheckCodeDo)
}

// ReleaseCCICCheckCodeDo 释放CCICCheckCodeDo
func ReleaseCCICCheckCodeDo(v *CCICCheckCodeDo) {
	v.EnterUrl = ""
	v.Enterable = false
	poolCCICCheckCodeDo.Put(v)
}
