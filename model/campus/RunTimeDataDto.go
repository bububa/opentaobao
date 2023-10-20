package campus

import (
	"sync"
)

// RunTimeDataDto 结构体
type RunTimeDataDto struct {
	// runTimeData
	RunTimeDatas []RunDataDto `json:"run_time_datas,omitempty" xml:"run_time_datas>run_data_dto,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 消息id
	MsgId string `json:"msg_id,omitempty" xml:"msg_id,omitempty"`
	// 事件时间
	EventTime int64 `json:"event_time,omitempty" xml:"event_time,omitempty"`
	// 是否是逻辑设备
	BeLogic bool `json:"be_logic,omitempty" xml:"be_logic,omitempty"`
}

var poolRunTimeDataDto = sync.Pool{
	New: func() any {
		return new(RunTimeDataDto)
	},
}

// GetRunTimeDataDto() 从对象池中获取RunTimeDataDto
func GetRunTimeDataDto() *RunTimeDataDto {
	return poolRunTimeDataDto.Get().(*RunTimeDataDto)
}

// ReleaseRunTimeDataDto 释放RunTimeDataDto
func ReleaseRunTimeDataDto(v *RunTimeDataDto) {
	v.RunTimeDatas = v.RunTimeDatas[:0]
	v.DeviceId = ""
	v.MsgId = ""
	v.EventTime = 0
	v.BeLogic = false
	poolRunTimeDataDto.Put(v)
}
