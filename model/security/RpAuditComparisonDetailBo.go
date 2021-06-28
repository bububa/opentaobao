package security

// RpAuditComparisonDetailBo 
/* model for simplify = false
type RpAuditComparisonDetailBo struct {

    // 分数
    
    Score   string `json:"score,omitempty"`
    

    // 对比结果
    
    Value  *struct {
        RpAuditValueBo  *RpAuditValueBo `json:"rp_audit_value_bo,omitempty"`
    } `json:"value,omitempty"`
    

    // 类型
    
    ResultType  *struct {
        RpAuditTypeBo  *RpAuditTypeBo `json:"rp_audit_type_bo,omitempty"`
    } `json:"result_type,omitempty"`
    

}
*/

// RpAuditComparisonDetailBo 
type RpAuditComparisonDetailBo struct {

    // 分数
    Score   string `json:"score,omitempty"`

    // 对比结果
    Value   *RpAuditValueBo `json:"value,omitempty"`

    // 类型
    ResultType   *RpAuditTypeBo `json:"result_type,omitempty"`

}
