package scbp

import (
	"sync"
)

// TopMatchedProductDto 结构体
type TopMatchedProductDto struct {
	// 是否强制绑定
	IsForceMatch string `json:"is_force_match,omitempty" xml:"is_force_match,omitempty"`
	// 是否设置优先推广
	IsPreferential string `json:"is_preferential,omitempty" xml:"is_preferential,omitempty"`
	// 产品标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// Y:产品推广中,N:产品未推广
	IsOffer string `json:"is_offer,omitempty" xml:"is_offer,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 推广评分星级取值[0-5]
	QsStar int64 `json:"qs_star,omitempty" xml:"qs_star,omitempty"`
}

var poolTopMatchedProductDto = sync.Pool{
	New: func() any {
		return new(TopMatchedProductDto)
	},
}

// GetTopMatchedProductDto() 从对象池中获取TopMatchedProductDto
func GetTopMatchedProductDto() *TopMatchedProductDto {
	return poolTopMatchedProductDto.Get().(*TopMatchedProductDto)
}

// ReleaseTopMatchedProductDto 释放TopMatchedProductDto
func ReleaseTopMatchedProductDto(v *TopMatchedProductDto) {
	v.IsForceMatch = ""
	v.IsPreferential = ""
	v.Subject = ""
	v.IsOffer = ""
	v.ProductId = 0
	v.QsStar = 0
	poolTopMatchedProductDto.Put(v)
}
