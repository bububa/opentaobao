package viapi

import (
	"sync"
)

// Context 结构体
type Context struct {
	// 检测文本命中的风险内容上下文内容
	Context string `json:"context,omitempty" xml:"context,omitempty"`
}

var poolContext = sync.Pool{
	New: func() any {
		return new(Context)
	},
}

// GetContext() 从对象池中获取Context
func GetContext() *Context {
	return poolContext.Get().(*Context)
}

// ReleaseContext 释放Context
func ReleaseContext(v *Context) {
	v.Context = ""
	poolContext.Put(v)
}
