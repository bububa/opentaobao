package security

import (
	"sync"
)

// TaskInfo 结构体
type TaskInfo struct {
	// 任务唯一标识
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 任务处理进度: 1-处理完成 (可立即通过对应的查询接口查询处理结果), 2-异步处理中(需等待app_info.callback_url接收到反向通知后再查询结果) ; 对于app_info.data_type为1目前都是异步处理，此字段返回2; 对于app_info.data_type为2目前都是同步处理，此字段返回1
	Progress int64 `json:"progress,omitempty" xml:"progress,omitempty"`
}

var poolTaskInfo = sync.Pool{
	New: func() any {
		return new(TaskInfo)
	},
}

// GetTaskInfo() 从对象池中获取TaskInfo
func GetTaskInfo() *TaskInfo {
	return poolTaskInfo.Get().(*TaskInfo)
}

// ReleaseTaskInfo 释放TaskInfo
func ReleaseTaskInfo(v *TaskInfo) {
	v.ItemId = ""
	v.Progress = 0
	poolTaskInfo.Put(v)
}
