package security

// RpAuditValue 结构体
type RpAuditValue struct {
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
