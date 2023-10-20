package alsc

import (
	"sync"
)

// RechargeAccountPropertyOpenInfo 结构体
type RechargeAccountPropertyOpenInfo struct {
	// 账户类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 账户值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolRechargeAccountPropertyOpenInfo = sync.Pool{
	New: func() any {
		return new(RechargeAccountPropertyOpenInfo)
	},
}

// GetRechargeAccountPropertyOpenInfo() 从对象池中获取RechargeAccountPropertyOpenInfo
func GetRechargeAccountPropertyOpenInfo() *RechargeAccountPropertyOpenInfo {
	return poolRechargeAccountPropertyOpenInfo.Get().(*RechargeAccountPropertyOpenInfo)
}

// ReleaseRechargeAccountPropertyOpenInfo 释放RechargeAccountPropertyOpenInfo
func ReleaseRechargeAccountPropertyOpenInfo(v *RechargeAccountPropertyOpenInfo) {
	v.Type = ""
	v.Value = 0
	poolRechargeAccountPropertyOpenInfo.Put(v)
}
