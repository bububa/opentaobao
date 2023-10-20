package security

import (
	"sync"
)

// RpAuditDetails 结构体
type RpAuditDetails struct {
	// comparisonDetail
	ComparisonDetailList []RpAuditComparisonDetail `json:"comparison_detail_list,omitempty" xml:"comparison_detail_list>rp_audit_comparison_detail,omitempty"`
	// materialDetail
	MaterialDetailList []RpAuditMaterialDetail `json:"material_detail_list,omitempty" xml:"material_detail_list>rp_audit_material_detail,omitempty"`
	// 材料审核结论
	MaterialDetails []RpAuditMaterialDetail `json:"material_details,omitempty" xml:"material_details>rp_audit_material_detail,omitempty"`
	// 结果信息
	ComparisonDetails []RpAuditComparisonDetail `json:"comparison_details,omitempty" xml:"comparison_details>rp_audit_comparison_detail,omitempty"`
	// auditFinishTime
	AuditFinishTime string `json:"audit_finish_time,omitempty" xml:"audit_finish_time,omitempty"`
	// reviewDeadline
	ReviewDeadline string `json:"review_deadline,omitempty" xml:"review_deadline,omitempty"`
}

var poolRpAuditDetails = sync.Pool{
	New: func() any {
		return new(RpAuditDetails)
	},
}

// GetRpAuditDetails() 从对象池中获取RpAuditDetails
func GetRpAuditDetails() *RpAuditDetails {
	return poolRpAuditDetails.Get().(*RpAuditDetails)
}

// ReleaseRpAuditDetails 释放RpAuditDetails
func ReleaseRpAuditDetails(v *RpAuditDetails) {
	v.ComparisonDetailList = v.ComparisonDetailList[:0]
	v.MaterialDetailList = v.MaterialDetailList[:0]
	v.MaterialDetails = v.MaterialDetails[:0]
	v.ComparisonDetails = v.ComparisonDetails[:0]
	v.AuditFinishTime = ""
	v.ReviewDeadline = ""
	poolRpAuditDetails.Put(v)
}
