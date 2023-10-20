package alihealth2

import (
	"sync"
)

// AlibabaAlihealthReserveDentalBindshopanditemResult 结构体
type AlibabaAlihealthReserveDentalBindshopanditemResult struct {
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthReserveDentalBindshopanditemResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthReserveDentalBindshopanditemResult)
	},
}

// GetAlibabaAlihealthReserveDentalBindshopanditemResult() 从对象池中获取AlibabaAlihealthReserveDentalBindshopanditemResult
func GetAlibabaAlihealthReserveDentalBindshopanditemResult() *AlibabaAlihealthReserveDentalBindshopanditemResult {
	return poolAlibabaAlihealthReserveDentalBindshopanditemResult.Get().(*AlibabaAlihealthReserveDentalBindshopanditemResult)
}

// ReleaseAlibabaAlihealthReserveDentalBindshopanditemResult 释放AlibabaAlihealthReserveDentalBindshopanditemResult
func ReleaseAlibabaAlihealthReserveDentalBindshopanditemResult(v *AlibabaAlihealthReserveDentalBindshopanditemResult) {
	v.ErrMessage = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAlihealthReserveDentalBindshopanditemResult.Put(v)
}
