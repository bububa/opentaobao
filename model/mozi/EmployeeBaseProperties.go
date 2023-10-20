package mozi

import (
	"sync"
)

// EmployeeBaseProperties 结构体
type EmployeeBaseProperties struct {
	// 手机区号
	CellPhoneAreaCode string `json:"cell_phone_area_code,omitempty" xml:"cell_phone_area_code,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 头像地址
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 机号码
	CellPhone string `json:"cell_phone,omitempty" xml:"cell_phone,omitempty"`
	// 工号
	EmployeeNumber string `json:"employee_number,omitempty" xml:"employee_number,omitempty"`
}

var poolEmployeeBaseProperties = sync.Pool{
	New: func() any {
		return new(EmployeeBaseProperties)
	},
}

// GetEmployeeBaseProperties() 从对象池中获取EmployeeBaseProperties
func GetEmployeeBaseProperties() *EmployeeBaseProperties {
	return poolEmployeeBaseProperties.Get().(*EmployeeBaseProperties)
}

// ReleaseEmployeeBaseProperties 释放EmployeeBaseProperties
func ReleaseEmployeeBaseProperties(v *EmployeeBaseProperties) {
	v.CellPhoneAreaCode = ""
	v.NickName = ""
	v.Avatar = ""
	v.Email = ""
	v.CellPhone = ""
	v.EmployeeNumber = ""
	poolEmployeeBaseProperties.Put(v)
}
