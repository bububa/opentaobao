package alihealth2

// Diag 结构体
type Diag struct {
	// 诊断码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 诊断名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
