package alihealth2

import (
	"sync"
)

// Patient 结构体
type Patient struct {
	// 人群标签
	Labels []string `json:"labels,omitempty" xml:"labels>string,omitempty"`
	// 淘系用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用药人性别（0未知1男2女）
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 用药人体重-单位kg
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 用药人年龄-单位：天
	Age int64 `json:"age,omitempty" xml:"age,omitempty"`
	// 用药人身高-单位cm
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}

var poolPatient = sync.Pool{
	New: func() any {
		return new(Patient)
	},
}

// GetPatient() 从对象池中获取Patient
func GetPatient() *Patient {
	return poolPatient.Get().(*Patient)
}

// ReleasePatient 释放Patient
func ReleasePatient(v *Patient) {
	v.Labels = v.Labels[:0]
	v.UserId = ""
	v.Gender = 0
	v.Weight = 0
	v.Age = 0
	v.Height = 0
	poolPatient.Put(v)
}
