package axintrade

import (
	"sync"
)

// AxinPayImgUploadResDto 结构体
type AxinPayImgUploadResDto struct {
	// 图片id
	ImgId string `json:"img_id,omitempty" xml:"img_id,omitempty"`
}

var poolAxinPayImgUploadResDto = sync.Pool{
	New: func() any {
		return new(AxinPayImgUploadResDto)
	},
}

// GetAxinPayImgUploadResDto() 从对象池中获取AxinPayImgUploadResDto
func GetAxinPayImgUploadResDto() *AxinPayImgUploadResDto {
	return poolAxinPayImgUploadResDto.Get().(*AxinPayImgUploadResDto)
}

// ReleaseAxinPayImgUploadResDto 释放AxinPayImgUploadResDto
func ReleaseAxinPayImgUploadResDto(v *AxinPayImgUploadResDto) {
	v.ImgId = ""
	poolAxinPayImgUploadResDto.Put(v)
}
