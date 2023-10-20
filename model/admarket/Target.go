package admarket

import (
	"sync"
)

// Target 结构体
type Target struct {
	// 广告目标类型
	TargetType string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 广告目标值
	TargetValue string `json:"target_value,omitempty" xml:"target_value,omitempty"`
}

var poolTarget = sync.Pool{
	New: func() any {
		return new(Target)
	},
}

// GetTarget() 从对象池中获取Target
func GetTarget() *Target {
	return poolTarget.Get().(*Target)
}

// ReleaseTarget 释放Target
func ReleaseTarget(v *Target) {
	v.TargetType = ""
	v.TargetValue = ""
	poolTarget.Put(v)
}
