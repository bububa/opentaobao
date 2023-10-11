package simba

// BrandInfoVo 结构体
type BrandInfoVo struct {
	// 品牌ID
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
}
