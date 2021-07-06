package security

// RpAuditType 结构体
type RpAuditType struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
