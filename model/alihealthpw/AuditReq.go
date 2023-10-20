package alihealthpw

import (
	"sync"
)

// AuditReq 结构体
type AuditReq struct {
	// 审核时间
	AuditTime string `json:"audit_time,omitempty" xml:"audit_time,omitempty"`
	// 申请单唯一编码
	ApplyUniqueCode string `json:"apply_unique_code,omitempty" xml:"apply_unique_code,omitempty"`
	// 审核状态
	AuditStatus string `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 医院信息
	HospitalAddress string `json:"hospital_address,omitempty" xml:"hospital_address,omitempty"`
	// 治疗时间
	TreatTime string `json:"treat_time,omitempty" xml:"treat_time,omitempty"`
	// 审核意见
	AuditRemark string `json:"audit_remark,omitempty" xml:"audit_remark,omitempty"`
}

var poolAuditReq = sync.Pool{
	New: func() any {
		return new(AuditReq)
	},
}

// GetAuditReq() 从对象池中获取AuditReq
func GetAuditReq() *AuditReq {
	return poolAuditReq.Get().(*AuditReq)
}

// ReleaseAuditReq 释放AuditReq
func ReleaseAuditReq(v *AuditReq) {
	v.AuditTime = ""
	v.ApplyUniqueCode = ""
	v.AuditStatus = ""
	v.HospitalAddress = ""
	v.TreatTime = ""
	v.AuditRemark = ""
	poolAuditReq.Put(v)
}
