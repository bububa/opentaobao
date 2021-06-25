package security

// RpAuditComparisonDetailBo 
type RpAuditComparisonDetailBo struct {

    // 分数
    Score   string `json:"score,omitempty"`

    // 对比结果
    Value   *RpAuditValueBo `json:"value,omitempty"`

    // 类型
    ResultType   *RpAuditTypeBo `json:"result_type,omitempty"`

}
