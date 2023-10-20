package alihealthmedical

import (
	"sync"
)

// UploadDataRequest 结构体
type UploadDataRequest struct {
	// 医生外部id
	DoctorUuid string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
	// 文件类型名
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
}

var poolUploadDataRequest = sync.Pool{
	New: func() any {
		return new(UploadDataRequest)
	},
}

// GetUploadDataRequest() 从对象池中获取UploadDataRequest
func GetUploadDataRequest() *UploadDataRequest {
	return poolUploadDataRequest.Get().(*UploadDataRequest)
}

// ReleaseUploadDataRequest 释放UploadDataRequest
func ReleaseUploadDataRequest(v *UploadDataRequest) {
	v.DoctorUuid = ""
	v.ContentType = ""
	poolUploadDataRequest.Put(v)
}
