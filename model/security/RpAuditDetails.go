package security

// RpAuditDetails 
/* model for simplify = false
type RpAuditDetails struct {

    // auditFinishTime
    
    AuditFinishTime   string `json:"audit_finish_time,omitempty"`
    

    // comparisonDetail
    
    ComparisonDetailList  struct {
        RpAuditComparisonDetail  []RpAuditComparisonDetail `json:"rp_audit_comparison_detail,omitempty"`
    } `json:"comparison_detail_list,omitempty"`
    

    // materialDetail
    
    MaterialDetailList  struct {
        RpAuditMaterialDetail  []RpAuditMaterialDetail `json:"rp_audit_material_detail,omitempty"`
    } `json:"material_detail_list,omitempty"`
    

    // reviewDeadline
    
    ReviewDeadline   string `json:"review_deadline,omitempty"`
    

    // 材料审核结论
    
    MaterialDetails  struct {
        RpAuditMaterialDetail  []RpAuditMaterialDetail `json:"rp_audit_material_detail,omitempty"`
    } `json:"material_details,omitempty"`
    

    // 结果信息
    
    ComparisonDetails  struct {
        RpAuditComparisonDetail  []RpAuditComparisonDetail `json:"rp_audit_comparison_detail,omitempty"`
    } `json:"comparison_details,omitempty"`
    

}
*/

// RpAuditDetails 
type RpAuditDetails struct {

    // auditFinishTime
    AuditFinishTime   string `json:"audit_finish_time,omitempty"`

    // comparisonDetail
    ComparisonDetailList   []RpAuditComparisonDetail `json:"comparison_detail_list,omitempty"`

    // materialDetail
    MaterialDetailList   []RpAuditMaterialDetail `json:"material_detail_list,omitempty"`

    // reviewDeadline
    ReviewDeadline   string `json:"review_deadline,omitempty"`

    // 材料审核结论
    MaterialDetails   []RpAuditMaterialDetail `json:"material_details,omitempty"`

    // 结果信息
    ComparisonDetails   []RpAuditComparisonDetail `json:"comparison_details,omitempty"`

}
