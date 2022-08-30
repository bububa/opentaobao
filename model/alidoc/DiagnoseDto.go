package alidoc

// DiagnoseDto 结构体
type DiagnoseDto struct {
	// icdCode
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// icdCode名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
