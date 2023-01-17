package alidoc

// DiseaseInfo 结构体
type DiseaseInfo struct {
	// 诊断名称
	DiseaseName string `json:"disease_name,omitempty" xml:"disease_name,omitempty"`
	// 诊断code
	DiseaseCode string `json:"disease_code,omitempty" xml:"disease_code,omitempty"`
}
