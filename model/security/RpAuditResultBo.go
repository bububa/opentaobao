package security

// RpAuditResultBo 
type RpAuditResultBo struct {

    // 复核状态被置位
    
    ReviewStatus   bool `json:"review_status,omitempty" xml:"review_status,omitempty"`
    

    // 如果 curGrade 和 requireGrade 相同则 gradeCertified 返回 true, 相反则返回false
    
    GradeCertified   bool `json:"grade_certified,omitempty" xml:"grade_certified,omitempty"`
    

    // 要求的实人级别
    
    RequireGrade   *RpGradeBo `json:"require_grade,omitempty" xml:"require_grade,omitempty"`
    

    // 当前的实人级别
    
    CurGrade   *RpGradeBo `json:"cur_grade,omitempty" xml:"cur_grade,omitempty"`
    

    // 当前的biz
    
    Biz   string `json:"biz,omitempty" xml:"biz,omitempty"`
    

    // 当前的审核状态
    
    AuditStatus   *RpAuditStatusBo `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
    

    // 结果信息
    
    RpAuditDetails   *RpAuditDetails `json:"rp_audit_details,omitempty" xml:"rp_audit_details,omitempty"`
    

    // 复核类型
    
    ReviewType   string `json:"review_type,omitempty" xml:"review_type,omitempty"`
    

}
