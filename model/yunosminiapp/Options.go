package yunosminiapp

import (
	"sync"
)

// Options 结构体
type Options struct {
	// 请求来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 活动步骤
	Step string `json:"step,omitempty" xml:"step,omitempty"`
}

var poolOptions = sync.Pool{
	New: func() any {
		return new(Options)
	},
}

// GetOptions() 从对象池中获取Options
func GetOptions() *Options {
	return poolOptions.Get().(*Options)
}

// ReleaseOptions 释放Options
func ReleaseOptions(v *Options) {
	v.Source = ""
	v.Step = ""
	poolOptions.Put(v)
}
