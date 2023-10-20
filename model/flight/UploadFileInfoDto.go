package flight

import (
	"sync"
)

// UploadFileInfoDto 结构体
type UploadFileInfoDto struct {
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 附件地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolUploadFileInfoDto = sync.Pool{
	New: func() any {
		return new(UploadFileInfoDto)
	},
}

// GetUploadFileInfoDto() 从对象池中获取UploadFileInfoDto
func GetUploadFileInfoDto() *UploadFileInfoDto {
	return poolUploadFileInfoDto.Get().(*UploadFileInfoDto)
}

// ReleaseUploadFileInfoDto 释放UploadFileInfoDto
func ReleaseUploadFileInfoDto(v *UploadFileInfoDto) {
	v.Name = ""
	v.Url = ""
	poolUploadFileInfoDto.Put(v)
}
