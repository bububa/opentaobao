package security

// RpStatusResultBo 
type RpStatusResultBo struct {
    // 对应的bizId
    Biz   string `json:"biz,omitempty" xml:"biz,omitempty"`
    // 如果 curGrade 和 requireGrade 相同则 gradeCertified 返回 true，相反则返回false
    GradeCertified   bool `json:"grade_certified,omitempty" xml:"grade_certified,omitempty"`
    // 复核状态被置位
    ReviewStatus   bool `json:"review_status,omitempty" xml:"review_status,omitempty"`
    // 实人等级
    CurGrade   *RpGradeBo `json:"cur_grade,omitempty" xml:"cur_grade,omitempty"`
    // 要求达到的实人等级
    RequireGrade   *RpGradeBo `json:"require_grade,omitempty" xml:"require_grade,omitempty"`
    // 审核详细信息
    RpAuditDetails   *RpAuditDetailsBos `json:"rp_audit_details,omitempty" xml:"rp_audit_details,omitempty"`
    // 审核状态信息
    RpAuditStatus   *RpAuditStatusBo `json:"rp_audit_status,omitempty" xml:"rp_audit_status,omitempty"`
    // 复核类型
    ReviewType   string `json:"review_type,omitempty" xml:"review_type,omitempty"`
}
