package alicom

import (
	"sync"
)

// DsfSupplierSpuVo 结构体
type DsfSupplierSpuVo struct {
	// 分组
	GroupList []Group `json:"group_list,omitempty" xml:"group_list>group,omitempty"`
	// 业务类型
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
}

var poolDsfSupplierSpuVo = sync.Pool{
	New: func() any {
		return new(DsfSupplierSpuVo)
	},
}

// GetDsfSupplierSpuVo() 从对象池中获取DsfSupplierSpuVo
func GetDsfSupplierSpuVo() *DsfSupplierSpuVo {
	return poolDsfSupplierSpuVo.Get().(*DsfSupplierSpuVo)
}

// ReleaseDsfSupplierSpuVo 释放DsfSupplierSpuVo
func ReleaseDsfSupplierSpuVo(v *DsfSupplierSpuVo) {
	v.GroupList = v.GroupList[:0]
	v.BusinessType = ""
	poolDsfSupplierSpuVo.Put(v)
}
