package alihouse

// FeatureDto 结构体
type FeatureDto struct {
	// 1
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 1
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 1
	Source string `json:"source,omitempty" xml:"source,omitempty"`
}
