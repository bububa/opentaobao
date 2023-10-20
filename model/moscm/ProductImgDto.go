package moscm

import (
	"sync"
)

// ProductImgDto 结构体
type ProductImgDto struct {
	// url地址（以http或https开头的绝对路径）
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 排序
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
}

var poolProductImgDto = sync.Pool{
	New: func() any {
		return new(ProductImgDto)
	},
}

// GetProductImgDto() 从对象池中获取ProductImgDto
func GetProductImgDto() *ProductImgDto {
	return poolProductImgDto.Get().(*ProductImgDto)
}

// ReleaseProductImgDto 释放ProductImgDto
func ReleaseProductImgDto(v *ProductImgDto) {
	v.Url = ""
	v.Id = 0
	v.Position = 0
	poolProductImgDto.Put(v)
}
