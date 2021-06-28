package security

// RpAuditResult 
/* model for simplify = false
type RpAuditResult struct {

    // auditStatus
    
    AuditStatus  *struct {
        RpAuditStatus  *RpAuditStatus `json:"rp_audit_status,omitempty"`
    } `json:"audit_status,omitempty"`
    

    // biz
    
    Biz   string `json:"biz,omitempty"`
    

    // curGrade
    
    CurGrade  *struct {
        RpGrade  *RpGrade `json:"rp_grade,omitempty"`
    } `json:"cur_grade,omitempty"`
    

    // gradeCertified
    
    GradeCertified   bool `json:"grade_certified,omitempty"`
    

    // requireGrade
    
    RequireGrade  *struct {
        RpGrade  *RpGrade `json:"rp_grade,omitempty"`
    } `json:"require_grade,omitempty"`
    

    // reviewStatus
    
    ReviewStatus   bool `json:"review_status,omitempty"`
    

    // reviewType
    
    ReviewType   string `json:"review_type,omitempty"`
    

    // rpAuditDetails
    
    RpAuditDetails  *struct {
        RpAuditDetails  *RpAuditDetails `json:"rp_audit_details,omitempty"`
    } `json:"rp_audit_details,omitempty"`
    

    // rpUserResult
    
    RpUserResult  *struct {
        RpUserResult  *RpUserResult `json:"rp_user_result,omitempty"`
    } `json:"rp_user_result,omitempty"`
    

    // 审核比对信息
    
    Results  struct {
        RpAuditComparisonDetailBo  []RpAuditComparisonDetailBo `json:"rp_audit_comparison_detail_bo,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

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
