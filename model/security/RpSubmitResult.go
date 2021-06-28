package security

// RpSubmitResult 
/* model for simplify = false
type RpSubmitResult struct {

    // extraInfo
    
    ExtraInfo   string `json:"extra_info,omitempty"`
    

    // rpAuditResult
    
    RpAuditResult  *struct {
        RpAuditResult  *RpAuditResult `json:"rp_audit_result,omitempty"`
    } `json:"rp_audit_result,omitempty"`
    

    // steps
    
    Steps  struct {
        RpStepItem  []RpStepItem `json:"rp_step_item,omitempty"`
    } `json:"steps,omitempty"`
    

    // uploadToken
    
    UploadToken  *struct {
        StsUploadToken  *StsUploadToken `json:"sts_upload_token,omitempty"`
    } `json:"upload_token,omitempty"`
    

}
*/

// RpSubmitResult 
type RpSubmitResult struct {

    // extraInfo
    ExtraInfo   string `json:"extra_info,omitempty"`

    // rpAuditResult
    RpAuditResult   *RpAuditResult `json:"rp_audit_result,omitempty"`

    // steps
    Steps   []RpStepItem `json:"steps,omitempty"`

    // uploadToken
    UploadToken   *StsUploadToken `json:"upload_token,omitempty"`

}
