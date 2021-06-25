package security

// RpAuditComparisonDetailBos 
type RpAuditComparisonDetailBos struct {

    // 比对分值
    Score   string `json:"score,omitempty"`

    // 比对结论
    Value   *RpAuditValueBo `json:"value,omitempty"`

    // 比对项
    ResultType   *RpAuditTypeBo `json:"result_type,omitempty"`

}
