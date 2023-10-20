package alitripmerchant

import (
	"sync"
)

// EnumVo 结构体
type EnumVo struct {
	// 枚举值名称
	EnumName string `json:"enum_name,omitempty" xml:"enum_name,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 枚举代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 其他备注
	Other string `json:"other,omitempty" xml:"other,omitempty"`
	// 排序
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
}

var poolEnumVo = sync.Pool{
	New: func() any {
		return new(EnumVo)
	},
}

// GetEnumVo() 从对象池中获取EnumVo
func GetEnumVo() *EnumVo {
	return poolEnumVo.Get().(*EnumVo)
}

// ReleaseEnumVo 释放EnumVo
func ReleaseEnumVo(v *EnumVo) {
	v.EnumName = ""
	v.Desc = ""
	v.Code = ""
	v.Other = ""
	v.Order = 0
	poolEnumVo.Put(v)
}
