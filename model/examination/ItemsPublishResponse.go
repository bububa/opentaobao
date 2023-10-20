package examination

import (
	"sync"
)

// ItemsPublishResponse 结构体
type ItemsPublishResponse struct {
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}

var poolItemsPublishResponse = sync.Pool{
	New: func() any {
		return new(ItemsPublishResponse)
	},
}

// GetItemsPublishResponse() 从对象池中获取ItemsPublishResponse
func GetItemsPublishResponse() *ItemsPublishResponse {
	return poolItemsPublishResponse.Get().(*ItemsPublishResponse)
}

// ReleaseItemsPublishResponse 释放ItemsPublishResponse
func ReleaseItemsPublishResponse(v *ItemsPublishResponse) {
	v.Msg = ""
	poolItemsPublishResponse.Put(v)
}
