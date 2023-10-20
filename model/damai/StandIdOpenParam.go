package damai

import (
	"sync"
)

// StandIdOpenParam 结构体
type StandIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 看台id
	StandId int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolStandIdOpenParam = sync.Pool{
	New: func() any {
		return new(StandIdOpenParam)
	},
}

// GetStandIdOpenParam() 从对象池中获取StandIdOpenParam
func GetStandIdOpenParam() *StandIdOpenParam {
	return poolStandIdOpenParam.Get().(*StandIdOpenParam)
}

// ReleaseStandIdOpenParam 释放StandIdOpenParam
func ReleaseStandIdOpenParam(v *StandIdOpenParam) {
	v.SupplierSecret = ""
	v.PerformId = 0
	v.StandId = 0
	v.SystemId = 0
	poolStandIdOpenParam.Put(v)
}
