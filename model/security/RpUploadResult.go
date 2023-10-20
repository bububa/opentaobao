package security

import (
	"sync"
)

// RpUploadResult 结构体
type RpUploadResult struct {
	// uploadId
	UploadId string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	// uploadStatus
	UploadStatus *RpUploadStatus `json:"upload_status,omitempty" xml:"upload_status,omitempty"`
}

var poolRpUploadResult = sync.Pool{
	New: func() any {
		return new(RpUploadResult)
	},
}

// GetRpUploadResult() 从对象池中获取RpUploadResult
func GetRpUploadResult() *RpUploadResult {
	return poolRpUploadResult.Get().(*RpUploadResult)
}

// ReleaseRpUploadResult 释放RpUploadResult
func ReleaseRpUploadResult(v *RpUploadResult) {
	v.UploadId = ""
	v.UploadStatus = nil
	poolRpUploadResult.Put(v)
}
