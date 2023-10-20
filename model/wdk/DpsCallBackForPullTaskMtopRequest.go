package wdk

import (
	"sync"
)

// DpsCallBackForPullTaskMtopRequest 结构体
type DpsCallBackForPullTaskMtopRequest struct {
	// 任务列表
	TaskCodeList []string `json:"task_code_list,omitempty" xml:"task_code_list>string,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolDpsCallBackForPullTaskMtopRequest = sync.Pool{
	New: func() any {
		return new(DpsCallBackForPullTaskMtopRequest)
	},
}

// GetDpsCallBackForPullTaskMtopRequest() 从对象池中获取DpsCallBackForPullTaskMtopRequest
func GetDpsCallBackForPullTaskMtopRequest() *DpsCallBackForPullTaskMtopRequest {
	return poolDpsCallBackForPullTaskMtopRequest.Get().(*DpsCallBackForPullTaskMtopRequest)
}

// ReleaseDpsCallBackForPullTaskMtopRequest 释放DpsCallBackForPullTaskMtopRequest
func ReleaseDpsCallBackForPullTaskMtopRequest(v *DpsCallBackForPullTaskMtopRequest) {
	v.TaskCodeList = v.TaskCodeList[:0]
	v.WarehouseCode = ""
	poolDpsCallBackForPullTaskMtopRequest.Put(v)
}
