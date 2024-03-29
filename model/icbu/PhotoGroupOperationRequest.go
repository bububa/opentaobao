package icbu

import (
	"sync"
)

// PhotoGroupOperationRequest 结构体
type PhotoGroupOperationRequest struct {
	// add操作中表示新增的分组名，rename操作中表示重命名后的分组名，delete操作不填
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// add表示新增分组，delete表示删除分组，rename表示重命名分组
	Operation string `json:"operation,omitempty" xml:"operation,omitempty"`
	// add操作中表示新增分组的父分组id，delete操作和rename操作表示要操作的分组id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

var poolPhotoGroupOperationRequest = sync.Pool{
	New: func() any {
		return new(PhotoGroupOperationRequest)
	},
}

// GetPhotoGroupOperationRequest() 从对象池中获取PhotoGroupOperationRequest
func GetPhotoGroupOperationRequest() *PhotoGroupOperationRequest {
	return poolPhotoGroupOperationRequest.Get().(*PhotoGroupOperationRequest)
}

// ReleasePhotoGroupOperationRequest 释放PhotoGroupOperationRequest
func ReleasePhotoGroupOperationRequest(v *PhotoGroupOperationRequest) {
	v.GroupName = ""
	v.Operation = ""
	v.GroupId = 0
	poolPhotoGroupOperationRequest.Put(v)
}
