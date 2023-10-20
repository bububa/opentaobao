package axintrade

import (
	"sync"
)

// AxinPayImgUploadDto 结构体
type AxinPayImgUploadDto struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 图片类型
	ImageType string `json:"image_type,omitempty" xml:"image_type,omitempty"`
}

var poolAxinPayImgUploadDto = sync.Pool{
	New: func() any {
		return new(AxinPayImgUploadDto)
	},
}

// GetAxinPayImgUploadDto() 从对象池中获取AxinPayImgUploadDto
func GetAxinPayImgUploadDto() *AxinPayImgUploadDto {
	return poolAxinPayImgUploadDto.Get().(*AxinPayImgUploadDto)
}

// ReleaseAxinPayImgUploadDto 释放AxinPayImgUploadDto
func ReleaseAxinPayImgUploadDto(v *AxinPayImgUploadDto) {
	v.FileName = ""
	v.ImageType = ""
	poolAxinPayImgUploadDto.Put(v)
}
