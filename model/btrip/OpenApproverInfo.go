package btrip

import (
	"sync"
)

// OpenApproverInfo 结构体
type OpenApproverInfo struct {
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 审批意见
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 审批人名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 审批人id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 审批人状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 审批人状态：0审批中 1已同意 2已拒绝 3已转交，4已取消 5已终止 6发起审批 7评论
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 审批人顺序
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
}

var poolOpenApproverInfo = sync.Pool{
	New: func() any {
		return new(OpenApproverInfo)
	},
}

// GetOpenApproverInfo() 从对象池中获取OpenApproverInfo
func GetOpenApproverInfo() *OpenApproverInfo {
	return poolOpenApproverInfo.Get().(*OpenApproverInfo)
}

// ReleaseOpenApproverInfo 释放OpenApproverInfo
func ReleaseOpenApproverInfo(v *OpenApproverInfo) {
	v.OperateTime = ""
	v.Note = ""
	v.UserName = ""
	v.UserId = ""
	v.StatusDesc = ""
	v.Status = 0
	v.Order = 0
	poolOpenApproverInfo.Put(v)
}
