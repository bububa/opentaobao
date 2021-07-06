package wlb

// PackageMaterialList 结构体
type PackageMaterialList struct {
	// 包材
	MaterialType string `json:"material_type,omitempty" xml:"material_type,omitempty"`
	// 数量
	MaterialQuantity int64 `json:"material_quantity,omitempty" xml:"material_quantity,omitempty"`
}
