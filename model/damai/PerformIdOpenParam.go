package damai

import (
	"sync"
)

// PerformIdOpenParam 结构体
type PerformIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolPerformIdOpenParam = sync.Pool{
	New: func() any {
		return new(PerformIdOpenParam)
	},
}

// GetPerformIdOpenParam() 从对象池中获取PerformIdOpenParam
func GetPerformIdOpenParam() *PerformIdOpenParam {
	return poolPerformIdOpenParam.Get().(*PerformIdOpenParam)
}

// ReleasePerformIdOpenParam 释放PerformIdOpenParam
func ReleasePerformIdOpenParam(v *PerformIdOpenParam) {
	v.SupplierSecret = ""
	v.PerformId = 0
	v.SystemId = 0
	poolPerformIdOpenParam.Put(v)
}
