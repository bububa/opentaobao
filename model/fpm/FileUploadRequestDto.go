package fpm

import (
	"sync"
)

// FileUploadRequestDto 结构体
type FileUploadRequestDto struct {
	// 应用代码(必填)
	AppCode string `json:"app_code,omitempty" xml:"app_code,omitempty"`
	// 签名字符串
	OuterSystemSignStr string `json:"outer_system_sign_str,omitempty" xml:"outer_system_sign_str,omitempty"`
	// 加密字符串
	OuterSystemEncryptStr string `json:"outer_system_encrypt_str,omitempty" xml:"outer_system_encrypt_str,omitempty"`
}

var poolFileUploadRequestDto = sync.Pool{
	New: func() any {
		return new(FileUploadRequestDto)
	},
}

// GetFileUploadRequestDto() 从对象池中获取FileUploadRequestDto
func GetFileUploadRequestDto() *FileUploadRequestDto {
	return poolFileUploadRequestDto.Get().(*FileUploadRequestDto)
}

// ReleaseFileUploadRequestDto 释放FileUploadRequestDto
func ReleaseFileUploadRequestDto(v *FileUploadRequestDto) {
	v.AppCode = ""
	v.OuterSystemSignStr = ""
	v.OuterSystemEncryptStr = ""
	poolFileUploadRequestDto.Put(v)
}
