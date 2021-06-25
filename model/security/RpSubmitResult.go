package security

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
