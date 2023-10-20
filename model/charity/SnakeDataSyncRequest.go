package charity

import (
	"sync"
)

// SnakeDataSyncRequest 结构体
type SnakeDataSyncRequest struct {
	// 数据
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// 请求Id 用于排查问题
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 数据类型
	Action string `json:"action,omitempty" xml:"action,omitempty"`
}

var poolSnakeDataSyncRequest = sync.Pool{
	New: func() any {
		return new(SnakeDataSyncRequest)
	},
}

// GetSnakeDataSyncRequest() 从对象池中获取SnakeDataSyncRequest
func GetSnakeDataSyncRequest() *SnakeDataSyncRequest {
	return poolSnakeDataSyncRequest.Get().(*SnakeDataSyncRequest)
}

// ReleaseSnakeDataSyncRequest 释放SnakeDataSyncRequest
func ReleaseSnakeDataSyncRequest(v *SnakeDataSyncRequest) {
	v.DataList = v.DataList[:0]
	v.EventId = ""
	v.Action = ""
	poolSnakeDataSyncRequest.Put(v)
}
