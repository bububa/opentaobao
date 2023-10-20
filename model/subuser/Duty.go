package subuser

import (
	"sync"
)

// Duty 结构体
type Duty struct {
	// 职务名称
	DutyName string `json:"duty_name,omitempty" xml:"duty_name,omitempty"`
	// 职务ID
	DutyId int64 `json:"duty_id,omitempty" xml:"duty_id,omitempty"`
	// 职务级别
	DutyLevel int64 `json:"duty_level,omitempty" xml:"duty_level,omitempty"`
}

var poolDuty = sync.Pool{
	New: func() any {
		return new(Duty)
	},
}

// GetDuty() 从对象池中获取Duty
func GetDuty() *Duty {
	return poolDuty.Get().(*Duty)
}

// ReleaseDuty 释放Duty
func ReleaseDuty(v *Duty) {
	v.DutyName = ""
	v.DutyId = 0
	v.DutyLevel = 0
	poolDuty.Put(v)
}
