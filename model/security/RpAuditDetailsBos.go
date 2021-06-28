package security

// RpAuditDetailsBos 
/* model for simplify = false
type RpAuditDetailsBos struct {

    // 预计审核结束时间
    
    AuditFinishTime   string `json:"audit_finish_time,omitempty"`
    

    // 复核截止时间
    
    ReviewDeadline   string `json:"review_deadline,omitempty"`
    

    // 比对审核结论
    
    ComparisonDetails  struct {
        RpAuditComparisonDetailBos  []RpAuditComparisonDetailBos `json:"rp_audit_comparison_detail_bos,omitempty"`
    } `json:"comparison_details,omitempty"`
    

    // 材料审核结论
    
    MaterialDetails  struct {
        RpAuditMaterialDetailBos  []RpAuditMaterialDetailBos `json:"rp_audit_material_detail_bos,omitempty"`
    } `json:"material_details,omitempty"`
    

}
*/

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
