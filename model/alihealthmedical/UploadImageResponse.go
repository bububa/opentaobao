package alihealthmedical

import (
	"sync"
)

// UploadImageResponse 结构体
type UploadImageResponse struct {
	// 文件key值
	ObjectKey string `json:"object_key,omitempty" xml:"object_key,omitempty"`
	// url
	FullUrl string `json:"full_url,omitempty" xml:"full_url,omitempty"`
}

var poolUploadImageResponse = sync.Pool{
	New: func() any {
		return new(UploadImageResponse)
	},
}

// GetUploadImageResponse() 从对象池中获取UploadImageResponse
func GetUploadImageResponse() *UploadImageResponse {
	return poolUploadImageResponse.Get().(*UploadImageResponse)
}

// ReleaseUploadImageResponse 释放UploadImageResponse
func ReleaseUploadImageResponse(v *UploadImageResponse) {
	v.ObjectKey = ""
	v.FullUrl = ""
	poolUploadImageResponse.Put(v)
}
