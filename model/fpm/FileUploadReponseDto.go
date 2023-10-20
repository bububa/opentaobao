package fpm

import (
	"sync"
)

// FileUploadReponseDto 结构体
type FileUploadReponseDto struct {
	// outerSystemSignStr
	OuterSystemSignStr string `json:"outer_system_sign_str,omitempty" xml:"outer_system_sign_str,omitempty"`
	// outerSystemEncryptStr
	OuterSystemEncryptStr string `json:"outer_system_encrypt_str,omitempty" xml:"outer_system_encrypt_str,omitempty"`
}

var poolFileUploadReponseDto = sync.Pool{
	New: func() any {
		return new(FileUploadReponseDto)
	},
}

// GetFileUploadReponseDto() 从对象池中获取FileUploadReponseDto
func GetFileUploadReponseDto() *FileUploadReponseDto {
	return poolFileUploadReponseDto.Get().(*FileUploadReponseDto)
}

// ReleaseFileUploadReponseDto 释放FileUploadReponseDto
func ReleaseFileUploadReponseDto(v *FileUploadReponseDto) {
	v.OuterSystemSignStr = ""
	v.OuterSystemEncryptStr = ""
	poolFileUploadReponseDto.Put(v)
}
