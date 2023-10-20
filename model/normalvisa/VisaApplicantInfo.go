package normalvisa

import (
	"sync"
)

// VisaApplicantInfo 结构体
type VisaApplicantInfo struct {
	// 申请人id
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 子订单号
	SubTcOrderId int64 `json:"sub_tc_order_id,omitempty" xml:"sub_tc_order_id,omitempty"`
}

var poolVisaApplicantInfo = sync.Pool{
	New: func() any {
		return new(VisaApplicantInfo)
	},
}

// GetVisaApplicantInfo() 从对象池中获取VisaApplicantInfo
func GetVisaApplicantInfo() *VisaApplicantInfo {
	return poolVisaApplicantInfo.Get().(*VisaApplicantInfo)
}

// ReleaseVisaApplicantInfo 释放VisaApplicantInfo
func ReleaseVisaApplicantInfo(v *VisaApplicantInfo) {
	v.ApplyId = ""
	v.SubTcOrderId = 0
	poolVisaApplicantInfo.Put(v)
}
