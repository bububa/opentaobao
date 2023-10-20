package jipiao

import (
	"sync"
)

// CorpInfo 结构体
type CorpInfo struct {
	// 申请人姓名
	ApplyName string `json:"apply_name,omitempty" xml:"apply_name,omitempty"`
	// 申请人工号
	ApplyNo string `json:"apply_no,omitempty" xml:"apply_no,omitempty"`
	// BPM的fromNO
	FormNo string `json:"form_no,omitempty" xml:"form_no,omitempty"`
	// 出差人工号
	TripPersonNo string `json:"trip_person_no,omitempty" xml:"trip_person_no,omitempty"`
	// 出差人姓名
	TripPersonName string `json:"trip_person_name,omitempty" xml:"trip_person_name,omitempty"`
	// 工作地点
	WorkSpace string `json:"work_space,omitempty" xml:"work_space,omitempty"`
	// 成本中心代码
	CostCenterCode string `json:"cost_center_code,omitempty" xml:"cost_center_code,omitempty"`
	// 成本中心
	CostCenter string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// 单据状态代码
	FormStatus string `json:"form_status,omitempty" xml:"form_status,omitempty"`
	// 单据状态描述
	ReceiptsStatus string `json:"receipts_status,omitempty" xml:"receipts_status,omitempty"`
	// 费用归属OU的CODE
	CostOu string `json:"cost_ou,omitempty" xml:"cost_ou,omitempty"`
	// 申请时间
	ApplyTime string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
	// 集团id
	CorprationId string `json:"corpration_id,omitempty" xml:"corpration_id,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
}

var poolCorpInfo = sync.Pool{
	New: func() any {
		return new(CorpInfo)
	},
}

// GetCorpInfo() 从对象池中获取CorpInfo
func GetCorpInfo() *CorpInfo {
	return poolCorpInfo.Get().(*CorpInfo)
}

// ReleaseCorpInfo 释放CorpInfo
func ReleaseCorpInfo(v *CorpInfo) {
	v.ApplyName = ""
	v.ApplyNo = ""
	v.FormNo = ""
	v.TripPersonNo = ""
	v.TripPersonName = ""
	v.WorkSpace = ""
	v.CostCenterCode = ""
	v.CostCenter = ""
	v.FormStatus = ""
	v.ReceiptsStatus = ""
	v.CostOu = ""
	v.ApplyTime = ""
	v.CorprationId = ""
	v.Extra = ""
	poolCorpInfo.Put(v)
}
