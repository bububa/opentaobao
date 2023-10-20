package nrpos

import (
	"sync"
)

// OfflineFileDto 结构体
type OfflineFileDto struct {
	// 文件名称
	FileKey string `json:"file_key,omitempty" xml:"file_key,omitempty"`
	// 文件下载地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolOfflineFileDto = sync.Pool{
	New: func() any {
		return new(OfflineFileDto)
	},
}

// GetOfflineFileDto() 从对象池中获取OfflineFileDto
func GetOfflineFileDto() *OfflineFileDto {
	return poolOfflineFileDto.Get().(*OfflineFileDto)
}

// ReleaseOfflineFileDto 释放OfflineFileDto
func ReleaseOfflineFileDto(v *OfflineFileDto) {
	v.FileKey = ""
	v.Url = ""
	poolOfflineFileDto.Put(v)
}
