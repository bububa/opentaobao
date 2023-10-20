package alihouse

import (
	"sync"
)

// CompanyBrandDto 结构体
type CompanyBrandDto struct {
	// 品牌logo
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 外部品牌ID
	OuterBrandId string `json:"outer_brand_id,omitempty" xml:"outer_brand_id,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 品牌标签
	BrandTags string `json:"brand_tags,omitempty" xml:"brand_tags,omitempty"`
	// 是否删除 1-是 0-否
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 品牌类别：0-未知 1-saas  2-普通中介 3-KA 4-开发商  5-交易服务
	BrandCategory int64 `json:"brand_category,omitempty" xml:"brand_category,omitempty"`
	// 是否为测试数据  1-是  0-否
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolCompanyBrandDto = sync.Pool{
	New: func() any {
		return new(CompanyBrandDto)
	},
}

// GetCompanyBrandDto() 从对象池中获取CompanyBrandDto
func GetCompanyBrandDto() *CompanyBrandDto {
	return poolCompanyBrandDto.Get().(*CompanyBrandDto)
}

// ReleaseCompanyBrandDto 释放CompanyBrandDto
func ReleaseCompanyBrandDto(v *CompanyBrandDto) {
	v.Logo = ""
	v.BrandName = ""
	v.OuterBrandId = ""
	v.Description = ""
	v.BrandTags = ""
	v.IsDeleted = 0
	v.BrandCategory = 0
	v.IsTest = 0
	poolCompanyBrandDto.Put(v)
}
