package alihealth2

import (
	"sync"
)

// AlibabaAlihealthReserveDentalModifyrestimeResult 结构体
type AlibabaAlihealthReserveDentalModifyrestimeResult struct {
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthReserveDentalModifyrestimeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthReserveDentalModifyrestimeResult)
	},
}

// GetAlibabaAlihealthReserveDentalModifyrestimeResult() 从对象池中获取AlibabaAlihealthReserveDentalModifyrestimeResult
func GetAlibabaAlihealthReserveDentalModifyrestimeResult() *AlibabaAlihealthReserveDentalModifyrestimeResult {
	return poolAlibabaAlihealthReserveDentalModifyrestimeResult.Get().(*AlibabaAlihealthReserveDentalModifyrestimeResult)
}

// ReleaseAlibabaAlihealthReserveDentalModifyrestimeResult 释放AlibabaAlihealthReserveDentalModifyrestimeResult
func ReleaseAlibabaAlihealthReserveDentalModifyrestimeResult(v *AlibabaAlihealthReserveDentalModifyrestimeResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Success = false
	poolAlibabaAlihealthReserveDentalModifyrestimeResult.Put(v)
}
