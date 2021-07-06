package security

// RpStepType 结构体
type RpStepType struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
