package security

import (
	"sync"
)

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

var poolRpSubmitResult = sync.Pool{
	New: func() any {
		return new(RpSubmitResult)
	},
}

// GetRpSubmitResult() 从对象池中获取RpSubmitResult
func GetRpSubmitResult() *RpSubmitResult {
	return poolRpSubmitResult.Get().(*RpSubmitResult)
}

// ReleaseRpSubmitResult 释放RpSubmitResult
func ReleaseRpSubmitResult(v *RpSubmitResult) {
	v.Steps = v.Steps[:0]
	v.ExtraInfo = ""
	v.RpAuditResult = nil
	v.UploadToken = nil
	poolRpSubmitResult.Put(v)
}
