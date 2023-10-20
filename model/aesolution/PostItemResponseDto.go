package aesolution

import (
	"sync"
)

// PostItemResponseDto 结构体
type PostItemResponseDto struct {
	// productId
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolPostItemResponseDto = sync.Pool{
	New: func() any {
		return new(PostItemResponseDto)
	},
}

// GetPostItemResponseDto() 从对象池中获取PostItemResponseDto
func GetPostItemResponseDto() *PostItemResponseDto {
	return poolPostItemResponseDto.Get().(*PostItemResponseDto)
}

// ReleasePostItemResponseDto 释放PostItemResponseDto
func ReleasePostItemResponseDto(v *PostItemResponseDto) {
	v.ProductId = 0
	poolPostItemResponseDto.Put(v)
}
