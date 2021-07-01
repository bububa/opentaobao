package security

// RpAuditComparisonDetail 结构体
type RpAuditComparisonDetail struct {
	// resultType
	ResultType *RpAuditType `json:"result_type,omitempty" xml:"result_type,omitempty"`
	// score
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// value
	Value *RpAuditValue `json:"value,omitempty" xml:"value,omitempty"`
}
