package btrip

import (
	"sync"
)

// InsuranceInfo 结构体
type InsuranceInfo struct {
	// 保险单号
	InsuranceNo string `json:"insurance_no,omitempty" xml:"insurance_no,omitempty"`
	// 保险类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 保险状态：1已出保 2已退保
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 保险金额
	Amount float64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolInsuranceInfo = sync.Pool{
	New: func() any {
		return new(InsuranceInfo)
	},
}

// GetInsuranceInfo() 从对象池中获取InsuranceInfo
func GetInsuranceInfo() *InsuranceInfo {
	return poolInsuranceInfo.Get().(*InsuranceInfo)
}

// ReleaseInsuranceInfo 释放InsuranceInfo
func ReleaseInsuranceInfo(v *InsuranceInfo) {
	v.InsuranceNo = ""
	v.Type = ""
	v.Status = 0
	v.Amount = 0
	poolInsuranceInfo.Put(v)
}
