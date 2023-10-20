package simba

import (
	"sync"
)

// TopServiceContext 结构体
type TopServiceContext struct {
	// api业务线编码,查询账户余额bizCode必须是universalBP
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}

var poolTopServiceContext = sync.Pool{
	New: func() any {
		return new(TopServiceContext)
	},
}

// GetTopServiceContext() 从对象池中获取TopServiceContext
func GetTopServiceContext() *TopServiceContext {
	return poolTopServiceContext.Get().(*TopServiceContext)
}

// ReleaseTopServiceContext 释放TopServiceContext
func ReleaseTopServiceContext(v *TopServiceContext) {
	v.BizCode = ""
	poolTopServiceContext.Put(v)
}
