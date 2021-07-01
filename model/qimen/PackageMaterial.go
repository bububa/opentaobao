package qimen

// PackageMaterial 结构体
type PackageMaterial struct {
	// 包材型号
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 包材的数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
