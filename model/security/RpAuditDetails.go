package security

// RpAuditDetails 
type RpAuditDetails struct {
    // auditFinishTime
    AuditFinishTime   string `json:"audit_finish_time,omitempty" xml:"audit_finish_time,omitempty"`
    // comparisonDetail
    ComparisonDetailList   []RpAuditComparisonDetail `json:"comparison_detail_list,omitempty" xml:"comparison_detail_list>rp_audit_comparison_detail,omitempty"`
    // materialDetail
    MaterialDetailList   []RpAuditMaterialDetail `json:"material_detail_list,omitempty" xml:"material_detail_list>rp_audit_material_detail,omitempty"`
    // reviewDeadline
    ReviewDeadline   string `json:"review_deadline,omitempty" xml:"review_deadline,omitempty"`
    // 材料审核结论
    MaterialDetails   []RpAuditMaterialDetail `json:"material_details,omitempty" xml:"material_details>rp_audit_material_detail,omitempty"`
    // 结果信息
    ComparisonDetails   []RpAuditComparisonDetail `json:"comparison_details,omitempty" xml:"comparison_details>rp_audit_comparison_detail,omitempty"`
}
