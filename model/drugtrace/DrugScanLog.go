package drugtrace

import (
	"sync"
)

// DrugScanLog 结构体
type DrugScanLog struct {
	// 包装规格
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 溯源码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 有效期
	ExpiryDate string `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
}

var poolDrugScanLog = sync.Pool{
	New: func() any {
		return new(DrugScanLog)
	},
}

// GetDrugScanLog() 从对象池中获取DrugScanLog
func GetDrugScanLog() *DrugScanLog {
	return poolDrugScanLog.Get().(*DrugScanLog)
}

// ReleaseDrugScanLog 释放DrugScanLog
func ReleaseDrugScanLog(v *DrugScanLog) {
	v.PkgSpec = ""
	v.Code = ""
	v.ExpiryDate = ""
	poolDrugScanLog.Put(v)
}
