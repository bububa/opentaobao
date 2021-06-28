package security

// RpStatusResultBo 
/* model for simplify = false
type RpStatusResultBo struct {

    // 对应的bizId
    
    Biz   string `json:"biz,omitempty"`
    

    // 如果 curGrade 和 requireGrade 相同则 gradeCertified 返回 true，相反则返回false
    
    GradeCertified   bool `json:"grade_certified,omitempty"`
    

    // 复核状态被置位
    
    ReviewStatus   bool `json:"review_status,omitempty"`
    

    // 实人等级
    
    CurGrade  *struct {
        RpGradeBo  *RpGradeBo `json:"rp_grade_bo,omitempty"`
    } `json:"cur_grade,omitempty"`
    

    // 要求达到的实人等级
    
    RequireGrade  *struct {
        RpGradeBo  *RpGradeBo `json:"rp_grade_bo,omitempty"`
    } `json:"require_grade,omitempty"`
    

    // 审核详细信息
    
    RpAuditDetails  *struct {
        RpAuditDetailsBos  *RpAuditDetailsBos `json:"rp_audit_details_bos,omitempty"`
    } `json:"rp_audit_details,omitempty"`
    

    // 审核状态信息
    
    RpAuditStatus  *struct {
        RpAuditStatusBo  *RpAuditStatusBo `json:"rp_audit_status_bo,omitempty"`
    } `json:"rp_audit_status,omitempty"`
    

    // 复核类型
    
    ReviewType   string `json:"review_type,omitempty"`
    

}
*/

// RpStatusResultBo 
type RpStatusResultBo struct {

    // 对应的bizId
    Biz   string `json:"biz,omitempty"`

    // 如果 curGrade 和 requireGrade 相同则 gradeCertified 返回 true，相反则返回false
    GradeCertified   bool `json:"grade_certified,omitempty"`

    // 复核状态被置位
    ReviewStatus   bool `json:"review_status,omitempty"`

    // 实人等级
    CurGrade   *RpGradeBo `json:"cur_grade,omitempty"`

    // 要求达到的实人等级
    RequireGrade   *RpGradeBo `json:"require_grade,omitempty"`

    // 审核详细信息
    RpAuditDetails   *RpAuditDetailsBos `json:"rp_audit_details,omitempty"`

    // 审核状态信息
    RpAuditStatus   *RpAuditStatusBo `json:"rp_audit_status,omitempty"`

    // 复核类型
    ReviewType   string `json:"review_type,omitempty"`

}
