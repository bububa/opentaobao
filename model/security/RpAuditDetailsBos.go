package security

// RpAuditDetailsBos 
type RpAuditDetailsBos struct {

    // 预计审核结束时间
    AuditFinishTime   string `json:"audit_finish_time,omitempty"`

    // 复核截止时间
    ReviewDeadline   string `json:"review_deadline,omitempty"`

    // 比对审核结论
    ComparisonDetails   []RpAuditComparisonDetailBos `json:"comparison_details,omitempty"`

    // 材料审核结论
    MaterialDetails   []RpAuditMaterialDetailBos `json:"material_details,omitempty"`

}
