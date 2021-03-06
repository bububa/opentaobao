package tmallgenie

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
