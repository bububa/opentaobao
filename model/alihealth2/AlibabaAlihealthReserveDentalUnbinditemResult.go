package alihealth2

import (
	"sync"
)

// AlibabaAlihealthReserveDentalUnbinditemResult 结构体
type AlibabaAlihealthReserveDentalUnbinditemResult struct {
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthReserveDentalUnbinditemResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthReserveDentalUnbinditemResult)
	},
}

// GetAlibabaAlihealthReserveDentalUnbinditemResult() 从对象池中获取AlibabaAlihealthReserveDentalUnbinditemResult
func GetAlibabaAlihealthReserveDentalUnbinditemResult() *AlibabaAlihealthReserveDentalUnbinditemResult {
	return poolAlibabaAlihealthReserveDentalUnbinditemResult.Get().(*AlibabaAlihealthReserveDentalUnbinditemResult)
}

// ReleaseAlibabaAlihealthReserveDentalUnbinditemResult 释放AlibabaAlihealthReserveDentalUnbinditemResult
func ReleaseAlibabaAlihealthReserveDentalUnbinditemResult(v *AlibabaAlihealthReserveDentalUnbinditemResult) {
	v.ErrMessage = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAlihealthReserveDentalUnbinditemResult.Put(v)
}
