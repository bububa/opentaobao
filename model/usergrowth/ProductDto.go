package usergrowth

import (
	"sync"
)

// ProductDto 结构体
type ProductDto struct {
	// 商品名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商品描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 图片链接
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 标签属性
	Tags string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 类别
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 唤端链接
	Deeplink string `json:"deeplink,omitempty" xml:"deeplink,omitempty"`
	// 商品点击链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 商品曝光链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 产品词/召回词
	ProductQuery string `json:"product_query,omitempty" xml:"product_query,omitempty"`
	// 实体词
	EntityWord string `json:"entity_word,omitempty" xml:"entity_word,omitempty"`
}

var poolProductDto = sync.Pool{
	New: func() any {
		return new(ProductDto)
	},
}

// GetProductDto() 从对象池中获取ProductDto
func GetProductDto() *ProductDto {
	return poolProductDto.Get().(*ProductDto)
}

// ReleaseProductDto 释放ProductDto
func ReleaseProductDto(v *ProductDto) {
	v.Name = ""
	v.Description = ""
	v.ImageUrl = ""
	v.Tags = ""
	v.Category = ""
	v.Deeplink = ""
	v.ClickUrl = ""
	v.ExposureUrl = ""
	v.ProductQuery = ""
	v.EntityWord = ""
	poolProductDto.Put(v)
}
