package scbp

import (
	"sync"
)

// RelevantProductDto 结构体
type RelevantProductDto struct {
	// 推荐理由
	ReasonList []string `json:"reason_list,omitempty" xml:"reason_list>string,omitempty"`
	// 品图片链接
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 星级
	QsStar int64 `json:"qs_star,omitempty" xml:"qs_star,omitempty"`
	// 品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 操作优推品得主键id
	AdGroupId int64 `json:"ad_group_id,omitempty" xml:"ad_group_id,omitempty"`
	// 是否优推
	IsPreferential bool `json:"is_preferential,omitempty" xml:"is_preferential,omitempty"`
}

var poolRelevantProductDto = sync.Pool{
	New: func() any {
		return new(RelevantProductDto)
	},
}

// GetRelevantProductDto() 从对象池中获取RelevantProductDto
func GetRelevantProductDto() *RelevantProductDto {
	return poolRelevantProductDto.Get().(*RelevantProductDto)
}

// ReleaseRelevantProductDto 释放RelevantProductDto
func ReleaseRelevantProductDto(v *RelevantProductDto) {
	v.ReasonList = v.ReasonList[:0]
	v.ImgUrl = ""
	v.ProductName = ""
	v.QsStar = 0
	v.ProductId = 0
	v.AdGroupId = 0
	v.IsPreferential = false
	poolRelevantProductDto.Put(v)
}
