package yunos

import (
	"sync"
)

// RdamContext 结构体
type RdamContext struct {
}

var poolRdamContext = sync.Pool{
	New: func() any {
		return new(RdamContext)
	},
}

// GetRdamContext() 从对象池中获取RdamContext
func GetRdamContext() *RdamContext {
	return poolRdamContext.Get().(*RdamContext)
}

// ReleaseRdamContext 释放RdamContext
func ReleaseRdamContext(v *RdamContext) {
	poolRdamContext.Put(v)
}
