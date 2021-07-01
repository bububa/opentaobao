package wms

// Packagemateriallist 结构体
type Packagemateriallist struct {
	// 包裹包材信息
	PackageMaterial *Packagematerial `json:"package_material,omitempty" xml:"package_material,omitempty"`
}
