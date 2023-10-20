package normalvisa

import (
	"sync"
)

// NormalVisaUserUnit 结构体
type NormalVisaUserUnit struct {
	// 姓
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// 护照号
	PassportNumber string `json:"passport_number,omitempty" xml:"passport_number,omitempty"`
	// 名
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
}

var poolNormalVisaUserUnit = sync.Pool{
	New: func() any {
		return new(NormalVisaUserUnit)
	},
}

// GetNormalVisaUserUnit() 从对象池中获取NormalVisaUserUnit
func GetNormalVisaUserUnit() *NormalVisaUserUnit {
	return poolNormalVisaUserUnit.Get().(*NormalVisaUserUnit)
}

// ReleaseNormalVisaUserUnit 释放NormalVisaUserUnit
func ReleaseNormalVisaUserUnit(v *NormalVisaUserUnit) {
	v.LastName = ""
	v.PassportNumber = ""
	v.FirstName = ""
	poolNormalVisaUserUnit.Put(v)
}
