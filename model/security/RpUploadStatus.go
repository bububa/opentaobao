package security

import (
	"sync"
)

// RpUploadStatus 结构体
type RpUploadStatus struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpUploadStatus = sync.Pool{
	New: func() any {
		return new(RpUploadStatus)
	},
}

// GetRpUploadStatus() 从对象池中获取RpUploadStatus
func GetRpUploadStatus() *RpUploadStatus {
	return poolRpUploadStatus.Get().(*RpUploadStatus)
}

// ReleaseRpUploadStatus 释放RpUploadStatus
func ReleaseRpUploadStatus(v *RpUploadStatus) {
	v.Name = ""
	v.Code = 0
	poolRpUploadStatus.Put(v)
}
