package tvupadmin

// BrandDo 结构体
type BrandDo struct {
	// brandId
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// brandName
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
}
