package tmallsc

import (
	"sync"
)

// ReserveAuditRequest 结构体
type ReserveAuditRequest struct {
	// 变更时间（日期）
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 变更时间（时间段）
	UpdateTimeRange string `json:"update_time_range,omitempty" xml:"update_time_range,omitempty"`
	// 审核操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 审核拒绝原因
	RejectReason string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	// 改派师傅手机号
	NewWorkerMobile string `json:"new_worker_mobile,omitempty" xml:"new_worker_mobile,omitempty"`
	// 审核工单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 审核状态
	State int64 `json:"state,omitempty" xml:"state,omitempty"`
}

var poolReserveAuditRequest = sync.Pool{
	New: func() any {
		return new(ReserveAuditRequest)
	},
}

// GetReserveAuditRequest() 从对象池中获取ReserveAuditRequest
func GetReserveAuditRequest() *ReserveAuditRequest {
	return poolReserveAuditRequest.Get().(*ReserveAuditRequest)
}

// ReleaseReserveAuditRequest 释放ReserveAuditRequest
func ReleaseReserveAuditRequest(v *ReserveAuditRequest) {
	v.UpdateTime = ""
	v.UpdateTimeRange = ""
	v.Operator = ""
	v.RejectReason = ""
	v.NewWorkerMobile = ""
	v.Id = 0
	v.State = 0
	poolReserveAuditRequest.Put(v)
}
