package security

// RpAuditStatus 结构体
type RpAuditStatus struct {
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
