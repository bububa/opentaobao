package usergrowth

import (
	"sync"
)

// TaskConfig 结构体
type TaskConfig struct {
	// 任务等级，可空
	Level string `json:"level,omitempty" xml:"level,omitempty"`
	// 通过安全校验，可空
	PassSafeCheck bool `json:"pass_safe_check,omitempty" xml:"pass_safe_check,omitempty"`
}

var poolTaskConfig = sync.Pool{
	New: func() any {
		return new(TaskConfig)
	},
}

// GetTaskConfig() 从对象池中获取TaskConfig
func GetTaskConfig() *TaskConfig {
	return poolTaskConfig.Get().(*TaskConfig)
}

// ReleaseTaskConfig 释放TaskConfig
func ReleaseTaskConfig(v *TaskConfig) {
	v.Level = ""
	v.PassSafeCheck = false
	poolTaskConfig.Put(v)
}
