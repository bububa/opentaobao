package cloudgame

import (
	"sync"
)

// OssDto 结构体
type OssDto struct {
	// oss bucket
	Bucket string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// 文件名, 不包含后缀
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// oss endpoint
	Endpoint string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// 上传文件自定义key, 如energy:活力页头像
	FilePurpose string `json:"file_purpose,omitempty" xml:"file_purpose,omitempty"`
	// oss完整filekey
	FileKey string `json:"file_key,omitempty" xml:"file_key,omitempty"`
	// oss文件后缀名
	FileExt string `json:"file_ext,omitempty" xml:"file_ext,omitempty"`
	// 文件类型, AUDIO,IMAGE, VIDEO, GIF
	FileType string `json:"file_type,omitempty" xml:"file_type,omitempty"`
}

var poolOssDto = sync.Pool{
	New: func() any {
		return new(OssDto)
	},
}

// GetOssDto() 从对象池中获取OssDto
func GetOssDto() *OssDto {
	return poolOssDto.Get().(*OssDto)
}

// ReleaseOssDto 释放OssDto
func ReleaseOssDto(v *OssDto) {
	v.Bucket = ""
	v.FileName = ""
	v.Endpoint = ""
	v.FilePurpose = ""
	v.FileKey = ""
	v.FileExt = ""
	v.FileType = ""
	poolOssDto.Put(v)
}
