package security

import (
	"sync"
)

// UploadStatus 结构体
type UploadStatus struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolUploadStatus = sync.Pool{
	New: func() any {
		return new(UploadStatus)
	},
}

// GetUploadStatus() 从对象池中获取UploadStatus
func GetUploadStatus() *UploadStatus {
	return poolUploadStatus.Get().(*UploadStatus)
}

// ReleaseUploadStatus 释放UploadStatus
func ReleaseUploadStatus(v *UploadStatus) {
	v.Name = ""
	v.Code = 0
	poolUploadStatus.Put(v)
}
