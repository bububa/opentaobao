package hotel

// StatInfo 结构体
type StatInfo struct {
	// 品牌类型信息
	BrandTypeList []BrandType `json:"brand_type_list,omitempty" xml:"brand_type_list>brand_type,omitempty"`
}
