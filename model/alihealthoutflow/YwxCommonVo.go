package alihealthoutflow

import (
	"sync"
)

// YwxCommonVo 结构体
type YwxCommonVo struct {
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolYwxCommonVo = sync.Pool{
	New: func() any {
		return new(YwxCommonVo)
	},
}

// GetYwxCommonVo() 从对象池中获取YwxCommonVo
func GetYwxCommonVo() *YwxCommonVo {
	return poolYwxCommonVo.Get().(*YwxCommonVo)
}

// ReleaseYwxCommonVo 释放YwxCommonVo
func ReleaseYwxCommonVo(v *YwxCommonVo) {
	v.Status = ""
	v.Message = ""
	poolYwxCommonVo.Put(v)
}
