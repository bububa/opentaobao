package security

// RpAuditResult 
type RpAuditResult struct {

    // auditStatus
    AuditStatus   *RpAuditStatus `json:"audit_status,omitempty"`

    // biz
    Biz   string `json:"biz,omitempty"`

    // curGrade
    CurGrade   *RpGrade `json:"cur_grade,omitempty"`

    // gradeCertified
    GradeCertified   bool `json:"grade_certified,omitempty"`

    // requireGrade
    RequireGrade   *RpGrade `json:"require_grade,omitempty"`

    // reviewStatus
    ReviewStatus   bool `json:"review_status,omitempty"`

    // reviewType
    ReviewType   string `json:"review_type,omitempty"`

    // rpAuditDetails
    RpAuditDetails   *RpAuditDetails `json:"rp_audit_details,omitempty"`

    // rpUserResult
    RpUserResult   *RpUserResult `json:"rp_user_result,omitempty"`

    // 审核比对信息
    Results   []RpAuditComparisonDetailBo `json:"results,omitempty"`

}
