package jstinteractive

import (
	"sync"
)

// InteractiveTaskQueryResponse 结构体
type InteractiveTaskQueryResponse struct {
	// 任务列表
	TaskList []InteractiveTask `json:"task_list,omitempty" xml:"task_list>interactive_task,omitempty"`
}

var poolInteractiveTaskQueryResponse = sync.Pool{
	New: func() any {
		return new(InteractiveTaskQueryResponse)
	},
}

// GetInteractiveTaskQueryResponse() 从对象池中获取InteractiveTaskQueryResponse
func GetInteractiveTaskQueryResponse() *InteractiveTaskQueryResponse {
	return poolInteractiveTaskQueryResponse.Get().(*InteractiveTaskQueryResponse)
}

// ReleaseInteractiveTaskQueryResponse 释放InteractiveTaskQueryResponse
func ReleaseInteractiveTaskQueryResponse(v *InteractiveTaskQueryResponse) {
	v.TaskList = v.TaskList[:0]
	poolInteractiveTaskQueryResponse.Put(v)
}
