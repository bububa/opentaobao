package alicom

import (
	"sync"
)

// GroupRequest 结构体
type GroupRequest struct {
	// SPU ID列表
	SpuIdList []int64 `json:"spu_id_list,omitempty" xml:"spu_id_list>int64,omitempty"`
	// 是否主接口
	Main bool `json:"main,omitempty" xml:"main,omitempty"`
}

var poolGroupRequest = sync.Pool{
	New: func() any {
		return new(GroupRequest)
	},
}

// GetGroupRequest() 从对象池中获取GroupRequest
func GetGroupRequest() *GroupRequest {
	return poolGroupRequest.Get().(*GroupRequest)
}

// ReleaseGroupRequest 释放GroupRequest
func ReleaseGroupRequest(v *GroupRequest) {
	v.SpuIdList = v.SpuIdList[:0]
	v.Main = false
	poolGroupRequest.Put(v)
}
