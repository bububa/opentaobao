package alihouse

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
	// 是否删除 1-是 0-否
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 品牌类别（0-二手房类别品牌 1-租房类别品牌）
	BrandCategory int64 `json:"brand_category,omitempty" xml:"brand_category,omitempty"`
	// 是否为测试数据  1-是  0-否
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
