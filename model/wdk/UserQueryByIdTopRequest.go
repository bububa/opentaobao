package wdk

import (
	"sync"
)

// UserQueryByIdTopRequest 结构体
type UserQueryByIdTopRequest struct {
	// 仓库code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 人员id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolUserQueryByIdTopRequest = sync.Pool{
	New: func() any {
		return new(UserQueryByIdTopRequest)
	},
}

// GetUserQueryByIdTopRequest() 从对象池中获取UserQueryByIdTopRequest
func GetUserQueryByIdTopRequest() *UserQueryByIdTopRequest {
	return poolUserQueryByIdTopRequest.Get().(*UserQueryByIdTopRequest)
}

// ReleaseUserQueryByIdTopRequest 释放UserQueryByIdTopRequest
func ReleaseUserQueryByIdTopRequest(v *UserQueryByIdTopRequest) {
	v.WarehouseCode = ""
	v.UserId = 0
	poolUserQueryByIdTopRequest.Put(v)
}
