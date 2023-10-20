package alihealthpw

import (
	"sync"
)

// ApplyDetailResp 结构体
type ApplyDetailResp struct {
	// 审核时间
	AuditTime string `json:"audit_time,omitempty" xml:"audit_time,omitempty"`
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 状态文案
	ApplyStatusDesc string `json:"apply_status_desc,omitempty" xml:"apply_status_desc,omitempty"`
	// 状态
	ApplyStatus string `json:"apply_status,omitempty" xml:"apply_status,omitempty"`
	// 申请类型
	ApplyType string `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 申请信息
	ApplyInfo *ApplyInfo `json:"apply_info,omitempty" xml:"apply_info,omitempty"`
}

var poolApplyDetailResp = sync.Pool{
	New: func() any {
		return new(ApplyDetailResp)
	},
}

// GetApplyDetailResp() 从对象池中获取ApplyDetailResp
func GetApplyDetailResp() *ApplyDetailResp {
	return poolApplyDetailResp.Get().(*ApplyDetailResp)
}

// ReleaseApplyDetailResp 释放ApplyDetailResp
func ReleaseApplyDetailResp(v *ApplyDetailResp) {
	v.AuditTime = ""
	v.ApplyTime = ""
	v.ApplyStatusDesc = ""
	v.ApplyStatus = ""
	v.ApplyType = ""
	v.ApplyInfo = nil
	poolApplyDetailResp.Put(v)
}
