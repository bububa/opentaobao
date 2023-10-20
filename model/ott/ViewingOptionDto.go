package ott

import (
	"sync"
)

// ViewingOptionDto 结构体
type ViewingOptionDto struct {
	// currency
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// license
	License string `json:"license,omitempty" xml:"license,omitempty"`
	// price
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// quality
	Quality string `json:"quality,omitempty" xml:"quality,omitempty"`
}

var poolViewingOptionDto = sync.Pool{
	New: func() any {
		return new(ViewingOptionDto)
	},
}

// GetViewingOptionDto() 从对象池中获取ViewingOptionDto
func GetViewingOptionDto() *ViewingOptionDto {
	return poolViewingOptionDto.Get().(*ViewingOptionDto)
}

// ReleaseViewingOptionDto 释放ViewingOptionDto
func ReleaseViewingOptionDto(v *ViewingOptionDto) {
	v.Currency = ""
	v.License = ""
	v.Price = ""
	v.Quality = ""
	poolViewingOptionDto.Put(v)
}
