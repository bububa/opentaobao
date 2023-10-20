package ascp

import (
	"sync"
)

// SpecialSignLineRules 结构体
type SpecialSignLineRules struct {
	// 到货线路表达规则（组）
	PromiseDesRules []PromiseDesRule `json:"promise_des_rules,omitempty" xml:"promise_des_rules>promise_des_rule,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
}

var poolSpecialSignLineRules = sync.Pool{
	New: func() any {
		return new(SpecialSignLineRules)
	},
}

// GetSpecialSignLineRules() 从对象池中获取SpecialSignLineRules
func GetSpecialSignLineRules() *SpecialSignLineRules {
	return poolSpecialSignLineRules.Get().(*SpecialSignLineRules)
}

// ReleaseSpecialSignLineRules 释放SpecialSignLineRules
func ReleaseSpecialSignLineRules(v *SpecialSignLineRules) {
	v.PromiseDesRules = v.PromiseDesRules[:0]
	v.WmsOwnerCode = ""
	poolSpecialSignLineRules.Put(v)
}
