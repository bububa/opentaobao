package security

// RpStartResult 结构体
type RpStartResult struct {
	// steps
	Steps []RpStepItem `json:"steps,omitempty" xml:"steps>rp_step_item,omitempty"`
	// biz
	Biz string `json:"biz,omitempty" xml:"biz,omitempty"`
	// extraInfo
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// source
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// uploadToken
	UploadToken *StsUploadToken `json:"upload_token,omitempty" xml:"upload_token,omitempty"`
}
