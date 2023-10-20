package lstvending

import (
	"sync"
)

// VendingImageDto 结构体
type VendingImageDto struct {
	// 图片唯一标识
	ImgPathId string `json:"img_path_id,omitempty" xml:"img_path_id,omitempty"`
	// 图片访问地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolVendingImageDto = sync.Pool{
	New: func() any {
		return new(VendingImageDto)
	},
}

// GetVendingImageDto() 从对象池中获取VendingImageDto
func GetVendingImageDto() *VendingImageDto {
	return poolVendingImageDto.Get().(*VendingImageDto)
}

// ReleaseVendingImageDto 释放VendingImageDto
func ReleaseVendingImageDto(v *VendingImageDto) {
	v.ImgPathId = ""
	v.Url = ""
	poolVendingImageDto.Put(v)
}
