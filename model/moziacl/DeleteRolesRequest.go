package moziacl

import (
	"sync"
)

// DeleteRolesRequest 结构体
type DeleteRolesRequest struct {
	// 要删除的角色code列表
	Names []string `json:"names,omitempty" xml:"names>string,omitempty"`
	// 操作主体
	PrincipalParam *BucUserPrincipalParam `json:"principal_param,omitempty" xml:"principal_param,omitempty"`
}

var poolDeleteRolesRequest = sync.Pool{
	New: func() any {
		return new(DeleteRolesRequest)
	},
}

// GetDeleteRolesRequest() 从对象池中获取DeleteRolesRequest
func GetDeleteRolesRequest() *DeleteRolesRequest {
	return poolDeleteRolesRequest.Get().(*DeleteRolesRequest)
}

// ReleaseDeleteRolesRequest 释放DeleteRolesRequest
func ReleaseDeleteRolesRequest(v *DeleteRolesRequest) {
	v.Names = v.Names[:0]
	v.PrincipalParam = nil
	poolDeleteRolesRequest.Put(v)
}
