package scbp

import (
	"sync"
)

// TopProductDto 结构体
type TopProductDto struct {
	// 产品推广状态，取值[disabled,enabled]
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 产品标题，最大长度256个字符
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolTopProductDto = sync.Pool{
	New: func() any {
		return new(TopProductDto)
	},
}

// GetTopProductDto() 从对象池中获取TopProductDto
func GetTopProductDto() *TopProductDto {
	return poolTopProductDto.Get().(*TopProductDto)
}

// ReleaseTopProductDto 释放TopProductDto
func ReleaseTopProductDto(v *TopProductDto) {
	v.Status = ""
	v.Subject = ""
	v.ProductId = 0
	poolTopProductDto.Put(v)
}
