package security

import (
	"sync"
)

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

var poolRpStartResult = sync.Pool{
	New: func() any {
		return new(RpStartResult)
	},
}

// GetRpStartResult() 从对象池中获取RpStartResult
func GetRpStartResult() *RpStartResult {
	return poolRpStartResult.Get().(*RpStartResult)
}

// ReleaseRpStartResult 释放RpStartResult
func ReleaseRpStartResult(v *RpStartResult) {
	v.Steps = v.Steps[:0]
	v.Biz = ""
	v.ExtraInfo = ""
	v.Source = ""
	v.UploadToken = nil
	poolRpStartResult.Put(v)
}
