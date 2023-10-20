package happytrip

import (
	"sync"
)

// AlibabaHappytripTaxiOrderComplainStruct 结构体
type AlibabaHappytripTaxiOrderComplainStruct struct {
	// 供应商工单号
	CaseId string `json:"case_id,omitempty" xml:"case_id,omitempty"`
}

var poolAlibabaHappytripTaxiOrderComplainStruct = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderComplainStruct)
	},
}

// GetAlibabaHappytripTaxiOrderComplainStruct() 从对象池中获取AlibabaHappytripTaxiOrderComplainStruct
func GetAlibabaHappytripTaxiOrderComplainStruct() *AlibabaHappytripTaxiOrderComplainStruct {
	return poolAlibabaHappytripTaxiOrderComplainStruct.Get().(*AlibabaHappytripTaxiOrderComplainStruct)
}

// ReleaseAlibabaHappytripTaxiOrderComplainStruct 释放AlibabaHappytripTaxiOrderComplainStruct
func ReleaseAlibabaHappytripTaxiOrderComplainStruct(v *AlibabaHappytripTaxiOrderComplainStruct) {
	v.CaseId = ""
	poolAlibabaHappytripTaxiOrderComplainStruct.Put(v)
}
