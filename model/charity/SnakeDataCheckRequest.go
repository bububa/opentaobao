package charity

import (
	"sync"
)

// SnakeDataCheckRequest 结构体
type SnakeDataCheckRequest struct {
	// 数据行为类型
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 开始时间-时间戳
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 数据总数
	DataCount int64 `json:"data_count,omitempty" xml:"data_count,omitempty"`
}

var poolSnakeDataCheckRequest = sync.Pool{
	New: func() any {
		return new(SnakeDataCheckRequest)
	},
}

// GetSnakeDataCheckRequest() 从对象池中获取SnakeDataCheckRequest
func GetSnakeDataCheckRequest() *SnakeDataCheckRequest {
	return poolSnakeDataCheckRequest.Get().(*SnakeDataCheckRequest)
}

// ReleaseSnakeDataCheckRequest 释放SnakeDataCheckRequest
func ReleaseSnakeDataCheckRequest(v *SnakeDataCheckRequest) {
	v.Action = ""
	v.StartTime = 0
	v.EndTime = 0
	v.DataCount = 0
	poolSnakeDataCheckRequest.Put(v)
}
