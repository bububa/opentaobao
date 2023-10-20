package tmallgenie

import (
	"sync"
)

// CreateAlarmParam 结构体
type CreateAlarmParam struct {
	// 设备uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 提醒方式
	Schedule *ScheduleDto `json:"schedule,omitempty" xml:"schedule,omitempty"`
	// 重复响铃次数
	TriggerRepeat int64 `json:"trigger_repeat,omitempty" xml:"trigger_repeat,omitempty"`
	// 铃声音量，取值范围：0-100
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
}

var poolCreateAlarmParam = sync.Pool{
	New: func() any {
		return new(CreateAlarmParam)
	},
}

// GetCreateAlarmParam() 从对象池中获取CreateAlarmParam
func GetCreateAlarmParam() *CreateAlarmParam {
	return poolCreateAlarmParam.Get().(*CreateAlarmParam)
}

// ReleaseCreateAlarmParam 释放CreateAlarmParam
func ReleaseCreateAlarmParam(v *CreateAlarmParam) {
	v.Uuid = ""
	v.Schedule = nil
	v.TriggerRepeat = 0
	v.Volume = 0
	poolCreateAlarmParam.Put(v)
}
