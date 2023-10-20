package campus

import (
	"sync"
)

// EventInfoApiDto 结构体
type EventInfoApiDto struct {
	// 事件时间
	EventTime string `json:"event_time,omitempty" xml:"event_time,omitempty"`
	// 规则实例code
	RuleInstanceCode string `json:"rule_instance_code,omitempty" xml:"rule_instance_code,omitempty"`
	// IBos报警事件状态上报
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 事件类型code
	EventTypeCode string `json:"event_type_code,omitempty" xml:"event_type_code,omitempty"`
	// 事件信息ID
	EventInfoId int64 `json:"event_info_id,omitempty" xml:"event_info_id,omitempty"`
}

var poolEventInfoApiDto = sync.Pool{
	New: func() any {
		return new(EventInfoApiDto)
	},
}

// GetEventInfoApiDto() 从对象池中获取EventInfoApiDto
func GetEventInfoApiDto() *EventInfoApiDto {
	return poolEventInfoApiDto.Get().(*EventInfoApiDto)
}

// ReleaseEventInfoApiDto 释放EventInfoApiDto
func ReleaseEventInfoApiDto(v *EventInfoApiDto) {
	v.EventTime = ""
	v.RuleInstanceCode = ""
	v.Data = ""
	v.EventTypeCode = ""
	v.EventInfoId = 0
	poolEventInfoApiDto.Put(v)
}
