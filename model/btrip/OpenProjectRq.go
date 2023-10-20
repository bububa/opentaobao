package btrip

import (
	"sync"
)

// OpenProjectRq 结构体
type OpenProjectRq struct {
	// 项目代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 第三方成本中心id
	ThirdPartCostCenterId string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// 第三方项目id
	ThirdPartId string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	// 第三方发票id
	ThirdPartInvoiceId string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenProjectRq = sync.Pool{
	New: func() any {
		return new(OpenProjectRq)
	},
}

// GetOpenProjectRq() 从对象池中获取OpenProjectRq
func GetOpenProjectRq() *OpenProjectRq {
	return poolOpenProjectRq.Get().(*OpenProjectRq)
}

// ReleaseOpenProjectRq 释放OpenProjectRq
func ReleaseOpenProjectRq(v *OpenProjectRq) {
	v.Code = ""
	v.CorpId = ""
	v.ProjectName = ""
	v.ThirdPartCostCenterId = ""
	v.ThirdPartId = ""
	v.ThirdPartInvoiceId = ""
	v.Version = 0
	poolOpenProjectRq.Put(v)
}
