package security

// RpAuditComparisonDetail 
/* model for simplify = false
type RpAuditComparisonDetail struct {

    // resultType
    
    ResultType  *struct {
        RpAuditType  *RpAuditType `json:"rp_audit_type,omitempty"`
    } `json:"result_type,omitempty"`
    

    // score
    
    Score   string `json:"score,omitempty"`
    

    // value
    
    Value  *struct {
        RpAuditValue  *RpAuditValue `json:"rp_audit_value,omitempty"`
    } `json:"value,omitempty"`
    

}
*/

// RpAuditComparisonDetail 
type RpAuditComparisonDetail struct {

    // resultType
    ResultType   *RpAuditType `json:"result_type,omitempty"`

    // score
    Score   string `json:"score,omitempty"`

    // value
    Value   *RpAuditValue `json:"value,omitempty"`

}
