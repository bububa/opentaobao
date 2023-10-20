package security

import (
	"sync"
)

// RpAuditDetailsBos 结构体
type RpAuditDetailsBos struct {
	// 比对审核结论
	ComparisonDetails []RpAuditComparisonDetailBos `json:"comparison_details,omitempty" xml:"comparison_details>rp_audit_comparison_detail_bos,omitempty"`
	// 材料审核结论
	MaterialDetails []RpAuditMaterialDetailBos `json:"material_details,omitempty" xml:"material_details>rp_audit_material_detail_bos,omitempty"`
	// 预计审核结束时间
	AuditFinishTime string `json:"audit_finish_time,omitempty" xml:"audit_finish_time,omitempty"`
	// 复核截止时间
	ReviewDeadline string `json:"review_deadline,omitempty" xml:"review_deadline,omitempty"`
}

var poolRpAuditDetailsBos = sync.Pool{
	New: func() any {
		return new(RpAuditDetailsBos)
	},
}

// GetRpAuditDetailsBos() 从对象池中获取RpAuditDetailsBos
func GetRpAuditDetailsBos() *RpAuditDetailsBos {
	return poolRpAuditDetailsBos.Get().(*RpAuditDetailsBos)
}

// ReleaseRpAuditDetailsBos 释放RpAuditDetailsBos
func ReleaseRpAuditDetailsBos(v *RpAuditDetailsBos) {
	v.ComparisonDetails = v.ComparisonDetails[:0]
	v.MaterialDetails = v.MaterialDetails[:0]
	v.AuditFinishTime = ""
	v.ReviewDeadline = ""
	poolRpAuditDetailsBos.Put(v)
}
