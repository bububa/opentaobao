package security

// RpSubmitResult 结构体
type RpSubmitResult struct {
	// steps
	Steps []RpStepItem `json:"steps,omitempty" xml:"steps>rp_step_item,omitempty"`
	// extraInfo
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// rpAuditResult
	RpAuditResult *RpAuditResult `json:"rp_audit_result,omitempty" xml:"rp_audit_result,omitempty"`
	// uploadToken
	UploadToken *StsUploadToken `json:"upload_token,omitempty" xml:"upload_token,omitempty"`
}
