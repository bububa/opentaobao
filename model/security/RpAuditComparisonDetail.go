package security

// RpAuditComparisonDetail 结构体
type RpAuditComparisonDetail struct {
	// score
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// resultType
	ResultType *RpAuditType `json:"result_type,omitempty" xml:"result_type,omitempty"`
	// value
	Value *RpAuditValue `json:"value,omitempty" xml:"value,omitempty"`
}
