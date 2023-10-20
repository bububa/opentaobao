package examination

import (
	"sync"
)

// StoreSpecialTagsResponse 结构体
type StoreSpecialTagsResponse struct {
	// 标签code
	TagCode string `json:"tag_code,omitempty" xml:"tag_code,omitempty"`
	// 标签name
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 优先级
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
}

var poolStoreSpecialTagsResponse = sync.Pool{
	New: func() any {
		return new(StoreSpecialTagsResponse)
	},
}

// GetStoreSpecialTagsResponse() 从对象池中获取StoreSpecialTagsResponse
func GetStoreSpecialTagsResponse() *StoreSpecialTagsResponse {
	return poolStoreSpecialTagsResponse.Get().(*StoreSpecialTagsResponse)
}

// ReleaseStoreSpecialTagsResponse 释放StoreSpecialTagsResponse
func ReleaseStoreSpecialTagsResponse(v *StoreSpecialTagsResponse) {
	v.TagCode = ""
	v.TagName = ""
	v.Priority = 0
	poolStoreSpecialTagsResponse.Put(v)
}
