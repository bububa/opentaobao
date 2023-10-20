package tbitem

import (
	"sync"
)

// SkuFeature 结构体
type SkuFeature struct {
	// colorHotNew
	ColorHotNew string `json:"color_hot_new,omitempty" xml:"color_hot_new,omitempty"`
	// colorMaterialImg
	ColorMaterialImg string `json:"color_material_img,omitempty" xml:"color_material_img,omitempty"`
	// colorValue
	ColorValue string `json:"color_value,omitempty" xml:"color_value,omitempty"`
	// colorMaterial
	ColorMaterial string `json:"color_material,omitempty" xml:"color_material,omitempty"`
	// colorSeries
	ColorSeries string `json:"color_series,omitempty" xml:"color_series,omitempty"`
}

var poolSkuFeature = sync.Pool{
	New: func() any {
		return new(SkuFeature)
	},
}

// GetSkuFeature() 从对象池中获取SkuFeature
func GetSkuFeature() *SkuFeature {
	return poolSkuFeature.Get().(*SkuFeature)
}

// ReleaseSkuFeature 释放SkuFeature
func ReleaseSkuFeature(v *SkuFeature) {
	v.ColorHotNew = ""
	v.ColorMaterialImg = ""
	v.ColorValue = ""
	v.ColorMaterial = ""
	v.ColorSeries = ""
	poolSkuFeature.Put(v)
}
