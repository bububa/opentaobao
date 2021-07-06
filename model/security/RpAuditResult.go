package security

// RpAuditResult 结构体
type RpAuditResult struct {
	// 审核比对信息
	Results []RpAuditComparisonDetailBo `json:"results,omitempty" xml:"results>rp_audit_comparison_detail_bo,omitempty"`
	// biz
	Biz string `json:"biz,omitempty" xml:"biz,omitempty"`
	// reviewType
	ReviewType string `json:"review_type,omitempty" xml:"review_type,omitempty"`
	// auditStatus
	AuditStatus *RpAuditStatus `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// curGrade
	CurGrade *RpGrade `json:"cur_grade,omitempty" xml:"cur_grade,omitempty"`
	// requireGrade
	RequireGrade *RpGrade `json:"require_grade,omitempty" xml:"require_grade,omitempty"`
	// rpAuditDetails
	RpAuditDetails *RpAuditDetails `json:"rp_audit_details,omitempty" xml:"rp_audit_details,omitempty"`
	// rpUserResult
	RpUserResult *RpUserResult `json:"rp_user_result,omitempty" xml:"rp_user_result,omitempty"`
	// gradeCertified
	GradeCertified bool `json:"grade_certified,omitempty" xml:"grade_certified,omitempty"`
	// reviewStatus
	ReviewStatus bool `json:"review_status,omitempty" xml:"review_status,omitempty"`
}
