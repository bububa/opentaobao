package fundplatform

import (
	"sync"
)

// CardFetchDetailDto 结构体
type CardFetchDetailDto struct {
	// 制卡模板号
	TemplateNo string `json:"template_no,omitempty" xml:"template_no,omitempty"`
	// 制卡数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 售价,单位为分。不填写则使用模板上配置的默认售价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolCardFetchDetailDto = sync.Pool{
	New: func() any {
		return new(CardFetchDetailDto)
	},
}

// GetCardFetchDetailDto() 从对象池中获取CardFetchDetailDto
func GetCardFetchDetailDto() *CardFetchDetailDto {
	return poolCardFetchDetailDto.Get().(*CardFetchDetailDto)
}

// ReleaseCardFetchDetailDto 释放CardFetchDetailDto
func ReleaseCardFetchDetailDto(v *CardFetchDetailDto) {
	v.TemplateNo = ""
	v.Num = 0
	v.Price = 0
	poolCardFetchDetailDto.Put(v)
}
