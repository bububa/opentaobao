package security

// RpAuditResultBo 
/* model for simplify = false
type RpAuditResultBo struct {

    // 复核状态被置位
    
    ReviewStatus   bool `json:"review_status,omitempty"`
    

    // 如果 curGrade 和 requireGrade 相同则 gradeCertified 返回 true, 相反则返回false
    
    GradeCertified   bool `json:"grade_certified,omitempty"`
    

    // 要求的实人级别
    
    RequireGrade  *struct {
        RpGradeBo  *RpGradeBo `json:"rp_grade_bo,omitempty"`
    } `json:"require_grade,omitempty"`
    

    // 当前的实人级别
    
    CurGrade  *struct {
        RpGradeBo  *RpGradeBo `json:"rp_grade_bo,omitempty"`
    } `json:"cur_grade,omitempty"`
    

    // 当前的biz
    
    Biz   string `json:"biz,omitempty"`
    

    // 当前的审核状态
    
    AuditStatus  *struct {
        RpAuditStatusBo  *RpAuditStatusBo `json:"rp_audit_status_bo,omitempty"`
    } `json:"audit_status,omitempty"`
    

    // 结果信息
    
    RpAuditDetails  *struct {
        RpAuditDetails  *RpAuditDetails `json:"rp_audit_details,omitempty"`
    } `json:"rp_audit_details,omitempty"`
    

    // 复核类型
    
    ReviewType   string `json:"review_type,omitempty"`
    

}
*/

// RpAuditResultBo 
type RpAuditResultBo struct {

    // 复核状态被置位
    ReviewStatus   bool `json:"review_status,omitempty"`

    // 如果 curGrade 和 requireGrade 相同则 gradeCertified 返回 true, 相反则返回false
    GradeCertified   bool `json:"grade_certified,omitempty"`

    // 要求的实人级别
    RequireGrade   *RpGradeBo `json:"require_grade,omitempty"`

    // 当前的实人级别
    CurGrade   *RpGradeBo `json:"cur_grade,omitempty"`

    // 当前的biz
    Biz   string `json:"biz,omitempty"`

    // 当前的审核状态
    AuditStatus   *RpAuditStatusBo `json:"audit_status,omitempty"`

    // 结果信息
    RpAuditDetails   *RpAuditDetails `json:"rp_audit_details,omitempty"`

    // 复核类型
    ReviewType   string `json:"review_type,omitempty"`

}
