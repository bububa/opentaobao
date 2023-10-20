package waybill

import (
	"sync"
)

// StandardTemplateResult 结构体
type StandardTemplateResult struct {
	// 该cp的所有标准模板
	StandardTemplates []StandardTemplateDo `json:"standard_templates,omitempty" xml:"standard_templates>standard_template_do,omitempty"`
	// cp编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}

var poolStandardTemplateResult = sync.Pool{
	New: func() any {
		return new(StandardTemplateResult)
	},
}

// GetStandardTemplateResult() 从对象池中获取StandardTemplateResult
func GetStandardTemplateResult() *StandardTemplateResult {
	return poolStandardTemplateResult.Get().(*StandardTemplateResult)
}

// ReleaseStandardTemplateResult 释放StandardTemplateResult
func ReleaseStandardTemplateResult(v *StandardTemplateResult) {
	v.StandardTemplates = v.StandardTemplates[:0]
	v.CpCode = ""
	poolStandardTemplateResult.Put(v)
}
