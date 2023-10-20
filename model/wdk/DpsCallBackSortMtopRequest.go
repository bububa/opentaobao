package wdk

import (
	"sync"
)

// DpsCallBackSortMtopRequest 结构体
type DpsCallBackSortMtopRequest struct {
	// 明细列表
	DetailRequestList []DpsCallBackSortDetailMtopRequest `json:"detail_request_list,omitempty" xml:"detail_request_list>dps_call_back_sort_detail_mtop_request,omitempty"`
	// 任务号
	TaskCode string `json:"task_code,omitempty" xml:"task_code,omitempty"`
	// 账号
	UserAccount string `json:"user_account,omitempty" xml:"user_account,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolDpsCallBackSortMtopRequest = sync.Pool{
	New: func() any {
		return new(DpsCallBackSortMtopRequest)
	},
}

// GetDpsCallBackSortMtopRequest() 从对象池中获取DpsCallBackSortMtopRequest
func GetDpsCallBackSortMtopRequest() *DpsCallBackSortMtopRequest {
	return poolDpsCallBackSortMtopRequest.Get().(*DpsCallBackSortMtopRequest)
}

// ReleaseDpsCallBackSortMtopRequest 释放DpsCallBackSortMtopRequest
func ReleaseDpsCallBackSortMtopRequest(v *DpsCallBackSortMtopRequest) {
	v.DetailRequestList = v.DetailRequestList[:0]
	v.TaskCode = ""
	v.UserAccount = ""
	v.WarehouseCode = ""
	poolDpsCallBackSortMtopRequest.Put(v)
}
