package jym

import (
	"sync"
)

// TaskSwitchReqDto 结构体
type TaskSwitchReqDto struct {
	// 游戏ID，0代表所有游戏
	GameId string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	// 原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 区服
	GameServer string `json:"game_server,omitempty" xml:"game_server,omitempty"`
	// 客户端ID, 0代表所有客户端
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// 开始时间，格式yyyy-MM-dd HH:mm:ss
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 结束时间，格式yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 任务类型，0-所有，1-验号，2-发布
	TaskType int64 `json:"task_type,omitempty" xml:"task_type,omitempty"`
	// 操作系统，0-所有，1-安卓，2-IOS
	OperationSystem int64 `json:"operation_system,omitempty" xml:"operation_system,omitempty"`
	// 规则状态，1-生效，2-失效
	RuleStatus int64 `json:"rule_status,omitempty" xml:"rule_status,omitempty"`
}

var poolTaskSwitchReqDto = sync.Pool{
	New: func() any {
		return new(TaskSwitchReqDto)
	},
}

// GetTaskSwitchReqDto() 从对象池中获取TaskSwitchReqDto
func GetTaskSwitchReqDto() *TaskSwitchReqDto {
	return poolTaskSwitchReqDto.Get().(*TaskSwitchReqDto)
}

// ReleaseTaskSwitchReqDto 释放TaskSwitchReqDto
func ReleaseTaskSwitchReqDto(v *TaskSwitchReqDto) {
	v.GameId = ""
	v.Reason = ""
	v.GameServer = ""
	v.ClientId = ""
	v.BeginTime = ""
	v.EndTime = ""
	v.TaskType = 0
	v.OperationSystem = 0
	v.RuleStatus = 0
	poolTaskSwitchReqDto.Put(v)
}
