package hotel

// Brand 结构体
type Brand struct {
	// 品牌code
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
}
