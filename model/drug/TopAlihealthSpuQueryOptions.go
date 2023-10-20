package drug

import (
	"sync"
)

// TopAlihealthSpuQueryOptions 结构体
type TopAlihealthSpuQueryOptions struct {
	// 查询选择器，是否查询说明书信息
	IncludeInstruction bool `json:"include_instruction,omitempty" xml:"include_instruction,omitempty"`
}

var poolTopAlihealthSpuQueryOptions = sync.Pool{
	New: func() any {
		return new(TopAlihealthSpuQueryOptions)
	},
}

// GetTopAlihealthSpuQueryOptions() 从对象池中获取TopAlihealthSpuQueryOptions
func GetTopAlihealthSpuQueryOptions() *TopAlihealthSpuQueryOptions {
	return poolTopAlihealthSpuQueryOptions.Get().(*TopAlihealthSpuQueryOptions)
}

// ReleaseTopAlihealthSpuQueryOptions 释放TopAlihealthSpuQueryOptions
func ReleaseTopAlihealthSpuQueryOptions(v *TopAlihealthSpuQueryOptions) {
	v.IncludeInstruction = false
	poolTopAlihealthSpuQueryOptions.Put(v)
}
