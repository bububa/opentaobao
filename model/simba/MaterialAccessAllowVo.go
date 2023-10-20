package simba

import (
	"sync"
)

// MaterialAccessAllowVo 结构体
type MaterialAccessAllowVo struct {
	// 准入不通过原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 物料id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 是否通过准入,true:是,false:否
	AccessAllowed bool `json:"access_allowed,omitempty" xml:"access_allowed,omitempty"`
}

var poolMaterialAccessAllowVo = sync.Pool{
	New: func() any {
		return new(MaterialAccessAllowVo)
	},
}

// GetMaterialAccessAllowVo() 从对象池中获取MaterialAccessAllowVo
func GetMaterialAccessAllowVo() *MaterialAccessAllowVo {
	return poolMaterialAccessAllowVo.Get().(*MaterialAccessAllowVo)
}

// ReleaseMaterialAccessAllowVo 释放MaterialAccessAllowVo
func ReleaseMaterialAccessAllowVo(v *MaterialAccessAllowVo) {
	v.Reason = ""
	v.MaterialId = 0
	v.AccessAllowed = false
	poolMaterialAccessAllowVo.Put(v)
}
