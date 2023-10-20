package wdk

import (
	"sync"
)

// DpsCallBackSortDetailMtopRequest 结构体
type DpsCallBackSortDetailMtopRequest struct {
	// 提交时间
	SubmitTime string `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	// 缺货数量
	LackNum string `json:"lack_num,omitempty" xml:"lack_num,omitempty"`
	// 提交数量
	SubmitNum string `json:"submit_num,omitempty" xml:"submit_num,omitempty"`
	// 容器号
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 明细操作人员code
	UserAccountCode string `json:"user_account_code,omitempty" xml:"user_account_code,omitempty"`
	// 明细id
	DetailId int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
}

var poolDpsCallBackSortDetailMtopRequest = sync.Pool{
	New: func() any {
		return new(DpsCallBackSortDetailMtopRequest)
	},
}

// GetDpsCallBackSortDetailMtopRequest() 从对象池中获取DpsCallBackSortDetailMtopRequest
func GetDpsCallBackSortDetailMtopRequest() *DpsCallBackSortDetailMtopRequest {
	return poolDpsCallBackSortDetailMtopRequest.Get().(*DpsCallBackSortDetailMtopRequest)
}

// ReleaseDpsCallBackSortDetailMtopRequest 释放DpsCallBackSortDetailMtopRequest
func ReleaseDpsCallBackSortDetailMtopRequest(v *DpsCallBackSortDetailMtopRequest) {
	v.SubmitTime = ""
	v.LackNum = ""
	v.SubmitNum = ""
	v.ContainerCode = ""
	v.UserAccountCode = ""
	v.DetailId = 0
	poolDpsCallBackSortDetailMtopRequest.Put(v)
}
