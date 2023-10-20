package alihealthalgo

import (
	"sync"
)

// SolutionResultTopSupport 结构体
type SolutionResultTopSupport struct {
	// 结果
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSolutionResultTopSupport = sync.Pool{
	New: func() any {
		return new(SolutionResultTopSupport)
	},
}

// GetSolutionResultTopSupport() 从对象池中获取SolutionResultTopSupport
func GetSolutionResultTopSupport() *SolutionResultTopSupport {
	return poolSolutionResultTopSupport.Get().(*SolutionResultTopSupport)
}

// ReleaseSolutionResultTopSupport 释放SolutionResultTopSupport
func ReleaseSolutionResultTopSupport(v *SolutionResultTopSupport) {
	v.DataList = v.DataList[:0]
	v.Message = ""
	v.Status = ""
	poolSolutionResultTopSupport.Put(v)
}
