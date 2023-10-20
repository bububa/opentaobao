package waybill

import (
	"sync"
)

// StandardTemplateDo 结构体
type StandardTemplateDo struct {
	// 模板名称
	StandardTemplateName string `json:"standard_template_name,omitempty" xml:"standard_template_name,omitempty"`
	// 模板url
	StandardTemplateUrl string `json:"standard_template_url,omitempty" xml:"standard_template_url,omitempty"`
	// 如果没有 brandCode,则为 default
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 模板id
	StandardTemplateId int64 `json:"standard_template_id,omitempty" xml:"standard_template_id,omitempty"`
	// 1 快递标准面单 ,2 快递三联面单, 3 快递便携式三联单, 4 快运标准面单, 5 快运三联面单, 6 快递一联单
	StandardWaybillType int64 `json:"standard_waybill_type,omitempty" xml:"standard_waybill_type,omitempty"`
}

var poolStandardTemplateDo = sync.Pool{
	New: func() any {
		return new(StandardTemplateDo)
	},
}

// GetStandardTemplateDo() 从对象池中获取StandardTemplateDo
func GetStandardTemplateDo() *StandardTemplateDo {
	return poolStandardTemplateDo.Get().(*StandardTemplateDo)
}

// ReleaseStandardTemplateDo 释放StandardTemplateDo
func ReleaseStandardTemplateDo(v *StandardTemplateDo) {
	v.StandardTemplateName = ""
	v.StandardTemplateUrl = ""
	v.BrandCode = ""
	v.StandardTemplateId = 0
	v.StandardWaybillType = 0
	poolStandardTemplateDo.Put(v)
}
