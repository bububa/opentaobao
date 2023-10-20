package security

import (
	"sync"
)

// RpAuditMaterialDetail 结构体
type RpAuditMaterialDetail struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// display
	Display string `json:"display,omitempty" xml:"display,omitempty"`
	// materialType
	MaterialType string `json:"material_type,omitempty" xml:"material_type,omitempty"`
	// suggestion
	Suggestion string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
	// text
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// type
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// intercept
	Intercept bool `json:"intercept,omitempty" xml:"intercept,omitempty"`
	// security
	Security bool `json:"security,omitempty" xml:"security,omitempty"`
}

var poolRpAuditMaterialDetail = sync.Pool{
	New: func() any {
		return new(RpAuditMaterialDetail)
	},
}

// GetRpAuditMaterialDetail() 从对象池中获取RpAuditMaterialDetail
func GetRpAuditMaterialDetail() *RpAuditMaterialDetail {
	return poolRpAuditMaterialDetail.Get().(*RpAuditMaterialDetail)
}

// ReleaseRpAuditMaterialDetail 释放RpAuditMaterialDetail
func ReleaseRpAuditMaterialDetail(v *RpAuditMaterialDetail) {
	v.Code = ""
	v.Display = ""
	v.MaterialType = ""
	v.Suggestion = ""
	v.Text = ""
	v.Type = ""
	v.Intercept = false
	v.Security = false
	poolRpAuditMaterialDetail.Put(v)
}
