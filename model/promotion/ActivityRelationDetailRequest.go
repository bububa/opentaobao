package promotion

import (
	"sync"
)

// ActivityRelationDetailRequest 结构体
type ActivityRelationDetailRequest struct {
	// 活动状态(VALID  ， DELETE)
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// ISV活动关联权益后获得的关联ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
}

var poolActivityRelationDetailRequest = sync.Pool{
	New: func() any {
		return new(ActivityRelationDetailRequest)
	},
}

// GetActivityRelationDetailRequest() 从对象池中获取ActivityRelationDetailRequest
func GetActivityRelationDetailRequest() *ActivityRelationDetailRequest {
	return poolActivityRelationDetailRequest.Get().(*ActivityRelationDetailRequest)
}

// ReleaseActivityRelationDetailRequest 释放ActivityRelationDetailRequest
func ReleaseActivityRelationDetailRequest(v *ActivityRelationDetailRequest) {
	v.Status = ""
	v.RelationId = 0
	poolActivityRelationDetailRequest.Put(v)
}
