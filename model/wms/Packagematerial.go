package wms

// Packagematerial 结构体
type Packagematerial struct {
	// 淘宝指定的包材型号
	MaterialType string `json:"material_type,omitempty" xml:"material_type,omitempty"`
	// 包材的数量
	MaterialQuantity string `json:"material_quantity,omitempty" xml:"material_quantity,omitempty"`
}
