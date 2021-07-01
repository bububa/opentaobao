package alidoc

// RxDiagnoseTopDto 结构体
type RxDiagnoseTopDto struct {
	// icdCode名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// icdCode
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
