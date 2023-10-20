package cloudgame

import (
	"sync"
)

// DeleteAppCommandResult 结构体
type DeleteAppCommandResult struct {
	// 游戏删除任务ID
	TaskId string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// 游戏删除是否成功
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

var poolDeleteAppCommandResult = sync.Pool{
	New: func() any {
		return new(DeleteAppCommandResult)
	},
}

// GetDeleteAppCommandResult() 从对象池中获取DeleteAppCommandResult
func GetDeleteAppCommandResult() *DeleteAppCommandResult {
	return poolDeleteAppCommandResult.Get().(*DeleteAppCommandResult)
}

// ReleaseDeleteAppCommandResult 释放DeleteAppCommandResult
func ReleaseDeleteAppCommandResult(v *DeleteAppCommandResult) {
	v.TaskId = ""
	v.Succeeded = false
	poolDeleteAppCommandResult.Put(v)
}
