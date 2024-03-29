package moscm

import (
	"sync"
)

// CspuCriteria 结构体
type CspuCriteria struct {
	// 货号
	ArtNos []string `json:"art_nos,omitempty" xml:"art_nos>string,omitempty"`
	// 条码列表
	Barcodes []string `json:"barcodes,omitempty" xml:"barcodes>string,omitempty"`
	// 品牌id列表
	BrandIds []string `json:"brand_ids,omitempty" xml:"brand_ids>string,omitempty"`
	// 类目
	CatIds []string `json:"cat_ids,omitempty" xml:"cat_ids>string,omitempty"`
	// cspu编码
	CspuIds []string `json:"cspu_ids,omitempty" xml:"cspu_ids>string,omitempty"`
	// 属性值id列表
	PropertyValueIds []string `json:"property_value_ids,omitempty" xml:"property_value_ids>string,omitempty"`
	// spuId列表
	SpuIds []string `json:"spu_ids,omitempty" xml:"spu_ids>string,omitempty"`
	// 按创建时间查询结束条件
	CreateDateEnd string `json:"create_date_end,omitempty" xml:"create_date_end,omitempty"`
	// 按创建时间查询开始条件
	CreateDateStart string `json:"create_date_start,omitempty" xml:"create_date_start,omitempty"`
	// 模糊查询
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 款号
	StyleNo string `json:"style_no,omitempty" xml:"style_no,omitempty"`
}

var poolCspuCriteria = sync.Pool{
	New: func() any {
		return new(CspuCriteria)
	},
}

// GetCspuCriteria() 从对象池中获取CspuCriteria
func GetCspuCriteria() *CspuCriteria {
	return poolCspuCriteria.Get().(*CspuCriteria)
}

// ReleaseCspuCriteria 释放CspuCriteria
func ReleaseCspuCriteria(v *CspuCriteria) {
	v.ArtNos = v.ArtNos[:0]
	v.Barcodes = v.Barcodes[:0]
	v.BrandIds = v.BrandIds[:0]
	v.CatIds = v.CatIds[:0]
	v.CspuIds = v.CspuIds[:0]
	v.PropertyValueIds = v.PropertyValueIds[:0]
	v.SpuIds = v.SpuIds[:0]
	v.CreateDateEnd = ""
	v.CreateDateStart = ""
	v.Keyword = ""
	v.StyleNo = ""
	poolCspuCriteria.Put(v)
}
