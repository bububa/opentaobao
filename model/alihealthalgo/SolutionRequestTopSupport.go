package alihealthalgo

import (
	"sync"
)

// SolutionRequestTopSupport 结构体
type SolutionRequestTopSupport struct {
	// 业务应用标识
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 业务参数
	Params string `json:"params,omitempty" xml:"params,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolSolutionRequestTopSupport = sync.Pool{
	New: func() any {
		return new(SolutionRequestTopSupport)
	},
}

// GetSolutionRequestTopSupport() 从对象池中获取SolutionRequestTopSupport
func GetSolutionRequestTopSupport() *SolutionRequestTopSupport {
	return poolSolutionRequestTopSupport.Get().(*SolutionRequestTopSupport)
}

// ReleaseSolutionRequestTopSupport 释放SolutionRequestTopSupport
func ReleaseSolutionRequestTopSupport(v *SolutionRequestTopSupport) {
	v.AppId = ""
	v.Params = ""
	v.UserId = ""
	poolSolutionRequestTopSupport.Put(v)
}
