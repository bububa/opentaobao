package security

// RpAuditComparisonDetailBos 
type RpAuditComparisonDetailBos struct {

    // 比对分值
    
    Score   string `json:"score,omitempty" xml:"score,omitempty"`
    

    // 比对结论
    
    Value   *RpAuditValueBo `json:"value,omitempty" xml:"value,omitempty"`
    

    // 比对项
    
    ResultType   *RpAuditTypeBo `json:"result_type,omitempty" xml:"result_type,omitempty"`
    

}
