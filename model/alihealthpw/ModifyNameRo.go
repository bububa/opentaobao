package alihealthpw

import (
	"sync"
)

// ModifyNameRo 结构体
type ModifyNameRo struct {
	// 患者姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 用户唯一编码
	UserUniqueCode string `json:"user_unique_code,omitempty" xml:"user_unique_code,omitempty"`
	// 三方项目id
	ProjectThirdId string `json:"project_third_id,omitempty" xml:"project_third_id,omitempty"`
}

var poolModifyNameRo = sync.Pool{
	New: func() any {
		return new(ModifyNameRo)
	},
}

// GetModifyNameRo() 从对象池中获取ModifyNameRo
func GetModifyNameRo() *ModifyNameRo {
	return poolModifyNameRo.Get().(*ModifyNameRo)
}

// ReleaseModifyNameRo 释放ModifyNameRo
func ReleaseModifyNameRo(v *ModifyNameRo) {
	v.PatientName = ""
	v.ModifyTime = ""
	v.UserUniqueCode = ""
	v.ProjectThirdId = ""
	poolModifyNameRo.Put(v)
}
