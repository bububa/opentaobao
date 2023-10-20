package mos

import (
	"sync"
)

// CallDispatcherRespDo 结构体
type CallDispatcherRespDo struct {
	// shortId
	ShortId string `json:"short_id,omitempty" xml:"short_id,omitempty"`
	// gotCode
	GotCode string `json:"got_code,omitempty" xml:"got_code,omitempty"`
	// packageCode
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// fulfillPlanId
	FulfillPlanId int64 `json:"fulfill_plan_id,omitempty" xml:"fulfill_plan_id,omitempty"`
}

var poolCallDispatcherRespDo = sync.Pool{
	New: func() any {
		return new(CallDispatcherRespDo)
	},
}

// GetCallDispatcherRespDo() 从对象池中获取CallDispatcherRespDo
func GetCallDispatcherRespDo() *CallDispatcherRespDo {
	return poolCallDispatcherRespDo.Get().(*CallDispatcherRespDo)
}

// ReleaseCallDispatcherRespDo 释放CallDispatcherRespDo
func ReleaseCallDispatcherRespDo(v *CallDispatcherRespDo) {
	v.ShortId = ""
	v.GotCode = ""
	v.PackageCode = ""
	v.FulfillPlanId = 0
	poolCallDispatcherRespDo.Put(v)
}
