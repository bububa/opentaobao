package seaking

import (
	"sync"
)

// TitleRewriteDetailDto 结构体
type TitleRewriteDetailDto struct {
	// 目标语种
	TargetLang string `json:"target_lang,omitempty" xml:"target_lang,omitempty"`
	// 源语种
	SourceLang string `json:"source_lang,omitempty" xml:"source_lang,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品所在平台（ae/lazada）
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 商品图片
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 子任务编号，从1开始，必须连续
	Idx int64 `json:"idx,omitempty" xml:"idx,omitempty"`
}

var poolTitleRewriteDetailDto = sync.Pool{
	New: func() any {
		return new(TitleRewriteDetailDto)
	},
}

// GetTitleRewriteDetailDto() 从对象池中获取TitleRewriteDetailDto
func GetTitleRewriteDetailDto() *TitleRewriteDetailDto {
	return poolTitleRewriteDetailDto.Get().(*TitleRewriteDetailDto)
}

// ReleaseTitleRewriteDetailDto 释放TitleRewriteDetailDto
func ReleaseTitleRewriteDetailDto(v *TitleRewriteDetailDto) {
	v.TargetLang = ""
	v.SourceLang = ""
	v.CategoryName = ""
	v.Title = ""
	v.Platform = ""
	v.ImageUrl = ""
	v.CategoryId = 0
	v.ProductId = 0
	v.Idx = 0
	poolTitleRewriteDetailDto.Put(v)
}
