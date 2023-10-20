package legalsuit

import (
	"sync"
)

// FileDto 结构体
type FileDto struct {
	// 文件名称
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 文件类型
	FileType string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	// 预览地址
	PreviewUrl string `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
	// 下载地址
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	// 文件中心id
	FileId int64 `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

var poolFileDto = sync.Pool{
	New: func() any {
		return new(FileDto)
	},
}

// GetFileDto() 从对象池中获取FileDto
func GetFileDto() *FileDto {
	return poolFileDto.Get().(*FileDto)
}

// ReleaseFileDto 释放FileDto
func ReleaseFileDto(v *FileDto) {
	v.FileName = ""
	v.FileType = ""
	v.PreviewUrl = ""
	v.DownloadUrl = ""
	v.FileId = 0
	poolFileDto.Put(v)
}
