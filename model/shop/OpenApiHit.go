package shop

import (
	"sync"
)

// OpenApiHit 结构体
type OpenApiHit struct {
	// 店铺ID
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 店铺ID
	EntityId string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// 店铺类型
	EntityType string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
	// 店铺信息字段
	RecmAttrs string `json:"recm_attrs,omitempty" xml:"recm_attrs,omitempty"`
	// 固定值
	RuleId string `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	// 排序权重
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 优先级
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
}

var poolOpenApiHit = sync.Pool{
	New: func() any {
		return new(OpenApiHit)
	},
}

// GetOpenApiHit() 从对象池中获取OpenApiHit
func GetOpenApiHit() *OpenApiHit {
	return poolOpenApiHit.Get().(*OpenApiHit)
}

// ReleaseOpenApiHit 释放OpenApiHit
func ReleaseOpenApiHit(v *OpenApiHit) {
	v.BizId = ""
	v.Desc = ""
	v.EntityId = ""
	v.EntityType = ""
	v.RecmAttrs = ""
	v.RuleId = ""
	v.Weight = ""
	v.Priority = 0
	poolOpenApiHit.Put(v)
}
