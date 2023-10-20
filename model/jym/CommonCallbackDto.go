package jym

import (
	"sync"
)

// CommonCallbackDto 结构体
type CommonCallbackDto struct {
	// 采集任务状态
	TaskState string `json:"task_state,omitempty" xml:"task_state,omitempty"`
	// 采集任务状态描述
	TaskMsg string `json:"task_msg,omitempty" xml:"task_msg,omitempty"`
	// 采集任务结果码
	TaskCode string `json:"task_code,omitempty" xml:"task_code,omitempty"`
	// 业务id
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 采集开始时间（时间戳/秒）
	TaskBeginTime string `json:"task_begin_time,omitempty" xml:"task_begin_time,omitempty"`
	// 任务id
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 采集结果
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

var poolCommonCallbackDto = sync.Pool{
	New: func() any {
		return new(CommonCallbackDto)
	},
}

// GetCommonCallbackDto() 从对象池中获取CommonCallbackDto
func GetCommonCallbackDto() *CommonCallbackDto {
	return poolCommonCallbackDto.Get().(*CommonCallbackDto)
}

// ReleaseCommonCallbackDto 释放CommonCallbackDto
func ReleaseCommonCallbackDto(v *CommonCallbackDto) {
	v.TaskState = ""
	v.TaskMsg = ""
	v.TaskCode = ""
	v.BizId = ""
	v.TaskBeginTime = ""
	v.TaskId = ""
	v.Content = ""
	poolCommonCallbackDto.Put(v)
}
