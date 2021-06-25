package security

// RpAuditComparisonDetail 
type RpAuditComparisonDetail struct {

    // resultType
    ResultType   *RpAuditType `json:"result_type,omitempty"`

    // score
    Score   string `json:"score,omitempty"`

    // value
    Value   *RpAuditValue `json:"value,omitempty"`

}
