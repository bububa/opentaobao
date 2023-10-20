package damai

import (
	"sync"
)

// QueryProjectResult 结构体
type QueryProjectResult struct {
	// 项目信息
	Projects []ProjectDto `json:"projects,omitempty" xml:"projects>project_dto,omitempty"`
}

var poolQueryProjectResult = sync.Pool{
	New: func() any {
		return new(QueryProjectResult)
	},
}

// GetQueryProjectResult() 从对象池中获取QueryProjectResult
func GetQueryProjectResult() *QueryProjectResult {
	return poolQueryProjectResult.Get().(*QueryProjectResult)
}

// ReleaseQueryProjectResult 释放QueryProjectResult
func ReleaseQueryProjectResult(v *QueryProjectResult) {
	v.Projects = v.Projects[:0]
	poolQueryProjectResult.Put(v)
}
