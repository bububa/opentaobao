package pentraprism

import (
	"sync"
)

// OpenTaskInfoVo 结构体
type OpenTaskInfoVo struct {
	// 任务奖励对象
	Rewards []TaskRewardVo `json:"rewards,omitempty" xml:"rewards>task_reward_vo,omitempty"`
	// 扩展型任务子列表对象
	SubList []OpenTaskInfoVo `json:"sub_list,omitempty" xml:"sub_list>open_task_info_vo,omitempty"`
	// 任务信息token，用于回流输入
	FromToken string `json:"from_token,omitempty" xml:"from_token,omitempty"`
	// 任务状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务子类型
	TaskSubType string `json:"task_sub_type,omitempty" xml:"task_sub_type,omitempty"`
	// 任务类型
	TaskType string `json:"task_type,omitempty" xml:"task_type,omitempty"`
	// 做任务时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 任务分组ID
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 任务分组所在位置，用户再次排序
	GroupIndex int64 `json:"group_index,omitempty" xml:"group_index,omitempty"`
	// 投放ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 分组内位置
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// 任务进度对象
	Progress *TaskProgressVo `json:"progress,omitempty" xml:"progress,omitempty"`
	// 扩展类任务标记，标记是否为当前任务
	IsCurrent bool `json:"is_current,omitempty" xml:"is_current,omitempty"`
	// 是否今天完成的
	IsToday bool `json:"is_today,omitempty" xml:"is_today,omitempty"`
}

var poolOpenTaskInfoVo = sync.Pool{
	New: func() any {
		return new(OpenTaskInfoVo)
	},
}

// GetOpenTaskInfoVo() 从对象池中获取OpenTaskInfoVo
func GetOpenTaskInfoVo() *OpenTaskInfoVo {
	return poolOpenTaskInfoVo.Get().(*OpenTaskInfoVo)
}

// ReleaseOpenTaskInfoVo 释放OpenTaskInfoVo
func ReleaseOpenTaskInfoVo(v *OpenTaskInfoVo) {
	v.Rewards = v.Rewards[:0]
	v.SubList = v.SubList[:0]
	v.FromToken = ""
	v.Status = ""
	v.TaskSubType = ""
	v.TaskType = ""
	v.Time = ""
	v.GroupId = 0
	v.GroupIndex = 0
	v.Id = 0
	v.Index = 0
	v.Progress = nil
	v.IsCurrent = false
	v.IsToday = false
	poolOpenTaskInfoVo.Put(v)
}
