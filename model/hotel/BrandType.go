package hotel

// BrandType 结构体
type BrandType struct {
	// 品牌列表
	BrandList []Brand `json:"brand_list,omitempty" xml:"brand_list>brand,omitempty"`
	// 品牌类型
	TypeId string `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// 品牌类型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
}
