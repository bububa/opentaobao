package security

// RpAuditComparisonDetailBos 
/* model for simplify = false
type RpAuditComparisonDetailBos struct {

    // 比对分值
    
    Score   string `json:"score,omitempty"`
    

    // 比对结论
    
    Value  *struct {
        RpAuditValueBo  *RpAuditValueBo `json:"rp_audit_value_bo,omitempty"`
    } `json:"value,omitempty"`
    

    // 比对项
    
    ResultType  *struct {
        RpAuditTypeBo  *RpAuditTypeBo `json:"rp_audit_type_bo,omitempty"`
    } `json:"result_type,omitempty"`
    

}
*/

// RpAuditComparisonDetailBos 
type RpAuditComparisonDetailBos struct {

    // 比对分值
    Score   string `json:"score,omitempty"`

    // 比对结论
    Value   *RpAuditValueBo `json:"value,omitempty"`

    // 比对项
    ResultType   *RpAuditTypeBo `json:"result_type,omitempty"`

}
