package fpm

import (
	"sync"
)

// ImageUploadRequest 结构体
type ImageUploadRequest struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 文件字节数组Base64字符串
	FileByteBase64Str string `json:"file_byte_base64_str,omitempty" xml:"file_byte_base64_str,omitempty"`
	// 文件大小
	FileSize int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
}

var poolImageUploadRequest = sync.Pool{
	New: func() any {
		return new(ImageUploadRequest)
	},
}

// GetImageUploadRequest() 从对象池中获取ImageUploadRequest
func GetImageUploadRequest() *ImageUploadRequest {
	return poolImageUploadRequest.Get().(*ImageUploadRequest)
}

// ReleaseImageUploadRequest 释放ImageUploadRequest
func ReleaseImageUploadRequest(v *ImageUploadRequest) {
	v.FileName = ""
	v.FileByteBase64Str = ""
	v.FileSize = 0
	poolImageUploadRequest.Put(v)
}
