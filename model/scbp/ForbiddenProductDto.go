package scbp

import (
	"sync"
)

// ForbiddenProductDto 结构体
type ForbiddenProductDto struct {
	// 标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolForbiddenProductDto = sync.Pool{
	New: func() any {
		return new(ForbiddenProductDto)
	},
}

// GetForbiddenProductDto() 从对象池中获取ForbiddenProductDto
func GetForbiddenProductDto() *ForbiddenProductDto {
	return poolForbiddenProductDto.Get().(*ForbiddenProductDto)
}

// ReleaseForbiddenProductDto 释放ForbiddenProductDto
func ReleaseForbiddenProductDto(v *ForbiddenProductDto) {
	v.Subject = ""
	v.ImgUrl = ""
	v.ProductId = 0
	v.Status = 0
	poolForbiddenProductDto.Put(v)
}
